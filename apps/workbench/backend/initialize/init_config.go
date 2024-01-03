package initialize

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/deamgo/workbench/conf"
	"github.com/deamgo/workbench/pkg/logger"
)

var (
	globalConfig conf.Config
	configPath   string
)

func InitConfig() {

	configPath = "conf" + consts.Separate + "config.yaml"
	config, err := os.ReadFile(configPath)
	if err != nil {
		logger.Errorf("Cannot open config file: %s", err.Error())
	}

	if err = yaml.Unmarshal(config, &globalConfig); err != nil {
		logger.Fatalf("initialize config file failed, err: %s", err.Error())
	}

	logger.Info("Init config successfully! GlobalConfig: ", globalConfig)
}

func GetConfig() conf.Config {
	return globalConfig
}
