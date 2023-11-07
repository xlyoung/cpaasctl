// 文件路径: internal/config/config.go

package config

import (
	"errors"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
	"os"
)

type AppConfig struct {
	Version string `yaml:"version"`
	// 添加其他可能的字段
}

type Config struct {
	Cpaas struct {
		Base     string `yaml:"base"`
		Registry struct {
			URL string `yaml:"url"`
		} `yaml:"registry"`
	} `yaml:"cpaas"`
	App map[string]AppConfig `yaml:"app"`
}

var (
	cfg *Config // 保存配置信息的全局变量
)

func InitConfig(filePath string) error {
	viper.SetConfigFile(filePath)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// 检查配置文件是否为空
	if viper.ConfigFileUsed() == "" {
		return errors.New("config file is empty")
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	return nil
}

func GetConfig() *Config {
	return cfg
}

func SaveConfig(cfg *Config, filePath string) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
