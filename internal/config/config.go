package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var Admin *adminConfig
var Secret *secretConfig
var Service serviceConfig
var Database databaseConfig

type serviceConfig struct {
	Port int `yaml:"port"`
}

type adminConfig struct {
	Password *string `yaml:"password"`
}

type secretConfig struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

type databaseConfig struct {
	Uri string `yaml:"uri"`
}

func MustInit(configPaths ...string) {
	var configPath *string

	if len(configPaths) > 0 {
		lastConfigPath := configPaths[len(configPaths)-1]
		configPath = &lastConfigPath
	}

	if configPath == nil {
		envConfig := os.Getenv("CONFIG")
		if envConfig != "" {
			configPath = &envConfig
		}
	}

	if configPath == nil {
		defaultPath := "config.yaml"
		configPath = &defaultPath
	}

	// 读取 YAML 文件内容
	data, err := os.ReadFile(*configPath)
	if err != nil {
		panic(fmt.Errorf("failed to read configuration file：%v", err))
	}

	var config struct {
		Server   serviceConfig   `yaml:"service"`
		Database databaseConfig `yaml:"database"`
		Admin    *adminConfig   `yaml:"admin"`
		Secret   *secretConfig  `yaml:"secret"`
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(fmt.Errorf("unable to parse YAML data：%v", err))
	}

	Admin = config.Admin
	Secret = config.Secret
	Service = config.Server
	Database = config.Database
}
