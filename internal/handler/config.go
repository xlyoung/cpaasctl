// 文件路径: internal/cmdhandler/handler.go

package handler

import (
	"errors"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/compose"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/config"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/logger"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/pkg/utils"
	"os"
	"path/filepath"
	"time"
)

func ViewConfig(appName string) error {
	// 获取配置对象
	cfg := config.GetConfig()

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

func UpdateConfig(cfgFile, appName, configName, configValue string) error {
	// 获取配置对象
	cfg := config.GetConfig()

	// 查找特定的应用
	appConfig, exists := cfg.App[appName]
	if !exists {
		errMsg := "app not found: " + appName
		logger.Logger.Println(errMsg)
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

	// 这里你可以将更改后的配置写回文件
	if err := config.SaveConfig(cfg, cfgFile); err != nil {
		logger.Logger.Printf("error writing config to file: %v", err) // 修改为使用 Logger
		return err
	} else {
		logger.Logger.Infof("Updated config for %s %s is %s\n", appName, configName, configValue)
	}
	return nil
}

func ExportDockercomposeFile(cfgFile string) error {

	cfg := config.GetConfig()

	projectDir := cfg.Cpaas.Base // 从配置中获取项目目录

	// 目录下寻找docker-compose.yaml文件
	dockerComposeFile := filepath.Join(projectDir, "docker-compose.yml")

	envVars, err := utils.SetEnvironmentVariables(cfg)
	if err != nil {
		return err
	}

	project, err := compose.LoadAndInterpolateComposeFile(dockerComposeFile, envVars)

	if err != nil {
		return err
	}

	// Create the output directory if it doesn't exist
	outputDir := filepath.Join(projectDir, "output")
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	currentTime := time.Now().Format("2006-01-02-15-04-05")
	outputPath := filepath.Join(outputDir, "docker-compose-"+currentTime+".yml")

	err = compose.ExportComposeFile(project, outputPath)

	if err != nil {
		return err
	}
	logger.Logger.Infof("Export docker-compose file to %s\n", outputPath)
	return nil
}
