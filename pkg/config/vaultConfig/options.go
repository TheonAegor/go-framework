package vaultConfig

import (
	"github.com/TheonAegor/go-framework/pkg/config"
	"github.com/hashicorp/vault/api"
)

type configKey struct{}

func Config(cfg *api.Config) config.Option {
	return config.SetOption(configKey{}, cfg)
}

type tokenKey struct{}

func Token(token string) config.Option {
	return config.SetOption(tokenKey{}, token)
}

type addrKey struct{}

func Address(addr string) config.Option {
	return config.SetOption(addrKey{}, addr)
}

type namespacePathKey struct{}

func NamespacePath(namespacePath string) config.Option {
	return config.SetOption(namespacePathKey{}, namespacePath)
}

type appPathKey struct{}

func AppPath(appPath string) config.Option {
	return config.SetOption(appPathKey{}, appPath)
}

type roleIDKey struct{}

func RoleID(role string) config.Option {
	return config.SetOption(roleIDKey{}, role)
}

type secretIDKey struct{}

func SecretID(secret string) config.Option {
	return config.SetOption(secretIDKey{}, secret)
}
