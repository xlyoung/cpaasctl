// 文件路径: internal/config/config.go

package config

import (
	"io/ioutil"
)

type Config struct {
	Cpaas struct {
		Base     string `yaml:"base"`
		Registry struct {
			URL string `yaml:"url"`
		} `yaml:"registry"`
	} `yaml:"cpaas"`
	App map[string]struct {
		Version string `yaml:"version"`
	} `yaml:"app"`
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
