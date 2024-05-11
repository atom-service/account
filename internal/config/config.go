package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var Admin *adminConfig
var Service serviceConfig
var Secrets []*secretConfig
var Database databaseConfig

type serviceConfig struct {
	Port int `yaml:"port"`
}

type adminConfig struct {
	Password *string `yaml:"password"`
}

type secretConfig struct {
	Name  string `yaml:"name"`
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
		Admin    *adminConfig    `yaml:"admin"`
		Server   serviceConfig   `yaml:"service"`
		Secrets  []*secretConfig `yaml:"secrets"`
		Database databaseConfig  `yaml:"database"`
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(fmt.Errorf("unable to parse YAML data：%v", err))
	}

	Admin = config.Admin
	Service = config.Server
	Secrets = config.Secrets
	Database = config.Database
}
