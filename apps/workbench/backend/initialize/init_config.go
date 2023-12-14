package initialize

import (
	"github.com/deamgo/workbench/conf"
	"github.com/deamgo/workbench/pkg/consts"
	"github.com/deamgo/workbench/pkg/logger"

	"os"
	"path"
	"runtime"

	"gopkg.in/yaml.v3"
)

var (
	globalConfig conf.Config
	configPath   string
)

func getRootDir() string {

	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Dir(filename))

	return root
}

func InitConfig() {

	configPath = getRootDir() + consts.Separate + "conf" + consts.Separate + "config.yaml"
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
