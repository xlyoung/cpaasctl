// 文件路径: internal/cmdhandler/handler.go

package handler

import (
	"errors"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/config" // 根据你的项目路径调整此行
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/logger"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

func ViewConfig(filePath, appName string) error {
	// 从文件加载配置
	cfg, err := config.LoadConfig(filePath)
	if err != nil {
		logger.Logger.Printf("error loading config: %v", err) // 使用 Logger
		return err
	}

	// 查找特定的应用
	appConfig, exists := cfg.App[appName]
	if !exists {
		errMsg := "app not found: " + appName
		logger.Logger.Println(errMsg) // 使用 Logger
		return errors.New(errMsg)
	}

	// 打印应用配置
	logger.Logger.Printf("Configuration for %s:\n", appName)
	logger.Logger.Printf("- Version: %s\n", appConfig.Version)
	// 在这里添加更多配置细节，如果有的话

	return nil
}

func UpdateConfig(filePath, appName, configName, configValue string) error {
	// 从文件加载配置
	cfg, err := config.LoadConfig(filePath)
	if err != nil {
		logger.Logger.Printf("error loading config: %v", err) // 修改为使用 Logger
		return err
	}

	// 查找特定的应用
	appConfig, exists := cfg.App[appName]
	if !exists {
		errMsg := "app not found: " + appName
		logger.Logger.Println(errMsg) // 修改为使用 Logger
		return errors.New(errMsg)
	}

	// 根据 configName 更新相应的配置
	switch configName {
	case "version":
		appConfig.Version = configValue
	default:
		errMsg := "unknown config name: " + configName
		logger.Logger.Println(errMsg) // 修改为使用 Logger
		return errors.New(errMsg)
	}

	// 更新 app 配置
	cfg.App[appName] = appConfig

	// 将更改后的配置写回文件
	data, err := yaml.Marshal(cfg)
	if err != nil {
		logger.Logger.Printf("error marshalling config: %v", err) // 修改为使用 Logger
		return err
	}
	if err := ioutil.WriteFile(filePath, data, 0644); err != nil {
		logger.Logger.Printf("error writing config to file: %v", err) // 修改为使用 Logger
		return err
	}

	return nil
}
