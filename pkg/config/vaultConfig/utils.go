package vaultConfig

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/vault/api"
)

var ErrVersionNotFound = errors.New("vault engine version info not found")

func getKVInfo(cli *api.Client, path string) (string, vaultEngineVersionType, error) {
	engineVersion, err := getKVMount(cli, path)
	if err != nil {
		return "", 0, err
	}

	switch engineVersion {
	case vaultEngineVersionV1.String():
		return "", vaultEngineVersionV1, nil
	case vaultEngineVersionV2.String():
		return "data", vaultEngineVersionV2, nil
	default:
		return "", 0, fmt.Errorf("%w: %s", ErrVersionNotFound, path)
	}
}

type mountInfo struct {
	Data struct {
		Options struct {
			Version string `json:"version"`
		} `json:"options"`
	} `json:"data"`
}

func getKVMount(cli *api.Client, path string) (string, error) {
	rsp, err := cli.RawRequest(cli.NewRequest("GET", "/v1/sys/internal/ui/mounts/"+path))
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()

	info := new(mountInfo)

	if err = json.NewDecoder(rsp.Body).Decode(info); err != nil {
		return "", err
	}

	if info.Data.Options.Version == "" {
		return "1", nil
	}

	return info.Data.Options.Version, nil
}
