package vaultConfig

import (
	"context"
	"dario.cat/mergo"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/vault/api"
	"gitlab.wildberries.ru/ext-delivery/ext-delivery/youtrack-report/pkg/config"
	"gitlab.wildberries.ru/ext-delivery/ext-delivery/youtrack-report/pkg/util"
	"path/filepath"
	"strconv"
)

const DefaultStructTag = "vault"

var (
	ErrPathNotExist   = errors.New("path is not exist")
	ErrUnknownVersion = errors.New("unknown vault engine version")
)

type vaultEngineVersionType int64

func (v vaultEngineVersionType) String() string {
	return strconv.FormatInt(int64(v), 10)
}

const (
	vaultEngineVersionV1 vaultEngineVersionType = 1
	vaultEngineVersionV2 vaultEngineVersionType = 2
)

type vaultConfig struct {
	opts          config.Options
	cli           *api.Client
	namespacePath string
	appPath       string
	roleID        string
	secretID      string
}

func (c *vaultConfig) Options() config.Options {
	return c.opts
}

func (c *vaultConfig) Init(opts ...config.Option) error {
	for _, o := range opts {
		o(&c.opts)
	}

	if c.opts.Codec == nil {
		return config.ErrCodecMissing
	}

	var (
		cfg   = api.DefaultConfig()
		token string
	)

	if c.opts.Context != nil {
		if v, ok := c.opts.Context.Value(configKey{}).(*api.Config); ok {
			cfg = v
		}

		if v, ok := c.opts.Context.Value(addrKey{}).(string); ok {
			cfg.Address = v
		}

		if v, ok := c.opts.Context.Value(tokenKey{}).(string); ok {
			token = v
		}

		if v, ok := c.opts.Context.Value(namespacePathKey{}).(string); ok {
			c.namespacePath = v
		}

		if v, ok := c.opts.Context.Value(appPathKey{}).(string); ok {
			c.appPath = v
		}

		if v, ok := c.opts.Context.Value(roleIDKey{}).(string); ok {
			c.roleID = v
		}

		if v, ok := c.opts.Context.Value(secretIDKey{}).(string); ok {
			c.secretID = v
		}
	}

	cli, err := api.NewClient(cfg)
	if err != nil {
		return err
	}

	cli.SetToken(token)

	c.cli = cli

	return nil
}

func (c *vaultConfig) Load(ctx context.Context, opts ...config.LoadOption) error {
	if err := c.load(ctx); err != nil && !c.opts.AllowFail {
		return err
	}

	return nil
}

func (c *vaultConfig) load(ctx context.Context) error {
	logger := c.Options().Logger

	for _, fn := range c.opts.BeforeLoad {
		if err := fn(ctx, c); err != nil {
			return err
		}
	}

	if c.cli.Token() == "" && c.roleID == "" && c.secretID == "" {
		logger.Info(ctx, "Vault initialization is skipped. No tokens provided.")
		return nil
	}

	if len(c.cli.Token()) == 0 {
		rsp, err := c.cli.Logical().Write("auth/approle/login", map[string]interface{}{
			"role_id":   c.roleID,
			"secret_id": c.secretID,
		})
		if err != nil {
			return err
		}
		c.cli.SetToken(rsp.Auth.ClientToken)
	}

	pathAdd, version, err := getKVInfo(c.cli, c.namespacePath)
	if err != nil {
		return err
	}

	path := filepath.Clean(c.namespacePath + "/" + pathAdd + "/" + c.appPath)

	pair, err := c.cli.Logical().Read(path)
	if err != nil {
		return fmt.Errorf("vault path %s not found %v", path, err)
	} else if pair == nil || pair.Data == nil {
		return fmt.Errorf("vault path %s not found %v", path, ErrPathNotExist)
	}

	data := make([]byte, 0)
	switch version {
	case vaultEngineVersionV1:
		data, err = json.Marshal(pair.Data)
	case vaultEngineVersionV2:
		data, err = json.Marshal(pair.Data["data"])
	default:
		return ErrUnknownVersion
	}

	if err != nil {
		return err
	}

	src, err := util.Zero(c.opts.Struct)
	if err != nil {
		return err
	}

	err = c.opts.Codec.Unmarshal(data, src)
	if err != nil {
		return err
	}

	if err = mergo.Merge(c.opts.Struct, src, mergo.WithOverride, mergo.WithTypeCheck, mergo.WithAppendSlice); err != nil {
		return err
	}

	for _, fn := range c.opts.AfterLoad {
		if err = fn(ctx, c); err != nil {
			return err
		}
	}

	return nil
}

func (c *vaultConfig) String() string {
	return "vault"
}

func (c *vaultConfig) Name() string {
	return c.opts.Name
}

func NewConfig(opts ...config.Option) config.Config {
	options := config.NewOptions(opts...)

	if len(options.StructTag) == 0 {
		options.StructTag = DefaultStructTag
	}

	return &vaultConfig{opts: options}
}
