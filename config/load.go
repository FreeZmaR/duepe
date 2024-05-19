package config

import (
	"duepe/config/types"
	"duepe/web"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

const configFileName = "duepe-config.json"

var (
	ErrorConfigNotFound = errors.New("config not found")
	ErrorConfigInvalid  = errors.New("config invalid")
)

type AppConfig struct {
	Database   *types.Database   `json:"database"`
	HTTPServer *types.HTTPServer `json:"http_server"`
}

func LoadCMDWeb() (*types.CMDWeb, error) {
	path, err := os.Executable()
	if err != nil {
		return nil, err
	}

	path = strings.Join(strings.Split(path, "/")[:len(strings.Split(path, "/"))-1], "/")

	if !isExistsConfigFile(path) {
		if errCreate := createConfigExample(path); errCreate != nil {
			slog.Error("error create example config file: ", slog.String("err:", errCreate.Error()))

			return nil, errCreate
		}

		return types.NewDefaultCMDWeb(), nil
	}

	appCfg, err := parseConfigFile(path)
	if err != nil {
		return types.NewDefaultCMDWeb(), nil
	}

	return &types.CMDWeb{
		HTTPServer: appCfg.HTTPServer,
		Database:   appCfg.Database,
		WebFs:      &types.WebFs{FS: web.GetPublicFs()},
	}, nil
}

func isExistsConfigFile(path string) bool {
	if path[len(path)-1] != '/' {
		path += "/"
	}

	path += configFileName

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func createConfigExample(path string) error {
	if path[len(path)-1] != '/' {
		path += "/"
	}

	path += configFileName

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	example := &AppConfig{
		Database:   types.NewDefaultDatabase(),
		HTTPServer: types.NewDefaultHTTPServer(),
	}

	if err = json.NewEncoder(file).Encode(example); err != nil {
		slog.Error("error write example config file: ", slog.String("err:", err.Error()))

		return err
	}

	slog.Info("example config file created", path)

	return nil
}

func parseConfigFile(path string) (*AppConfig, error) {
	if path[len(path)-1] != '/' {
		path += "/"
	}

	path += configFileName

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &AppConfig{}

	if errDecode := json.NewDecoder(file).Decode(config); errDecode != nil {
		return nil, fmt.Errorf("%w: %s", ErrorConfigInvalid, errDecode.Error())
	}

	return config, nil
}
