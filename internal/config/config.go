// 文件路径: internal/config/config.go

package config

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
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

func LoadConfig(filePath string) (*Config, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
