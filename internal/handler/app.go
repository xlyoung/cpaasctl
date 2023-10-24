package handler

import (
	"fmt"
	"github.com/compose-spec/compose-go/types"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/compose"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/config"
	logger "gitlab.hycyg.com/paas-tools/cpaasctl/internal/logger"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/utils"
)

func StartApp(appName string) error {
	// TODO: 实现应用的启动逻辑

	//加载 配置文件引入变量
	cfg, err := config.LoadConfig("config/cpaas.yaml")
	if err != nil {
		return fmt.Errorf("error loading config: %w", err)
	}

	envVars, err := utils.SetEnvironmentVariables(cfg)

	if err != nil {
		return fmt.Errorf("error setting environment variables: %w", err)
	}
	// 例如，使用Docker SDK启动容器或发送请求到一个API
	composeConfig, err := compose.LoadComposeFile("./docker-compose.yml", envVars)

	if err != nil {
		logger.Logger.Error("Error loading docker-compose file: %s\n", err)
		return err
	}
	var serviceConfig *types.ServiceConfig
	for _, service := range composeConfig.Config.Services {
		if service.Name == appName {
			serviceConfig = &service
			break
		}
	}

	if serviceConfig == nil {
		logger.Logger.Error("No service found with name: %s\n", appName)
		return fmt.Errorf("no service found with name: %s", appName)
	}
	err = compose.StartApp(appName)
	if err != nil {
		return fmt.Errorf("error starting app: %w", err)
	}
	logger.Logger.Infof("Starting %s...\n", appName)
	// 模拟成功的情况
	return nil
}

func StopApp(appName string) error {
	// TODO: 实现应用的停止逻辑
	// 例如，使用Docker SDK停止容器或发送请求到一个API
	fmt.Printf("Stopping %s...\n", appName)
	// 模拟成功的情况
	return nil
}

func RestartApp(appName string) error {
	// TODO: 实现应用的重启逻辑
	// 例如，使用Docker SDK重启容器或发送请求到一个API
	fmt.Printf("Restarting %s...\n", appName)
	// 模拟成功的情况
	return nil
}

func GetAppStatus(appName string) error {
	// TODO: 实现获取应用状态的逻辑
	// 例如，查询Docker容器的状态或发送请求到一个API
	fmt.Printf("Status of %s:\n", appName)
	// 模拟输出状态
	fmt.Println("Running")

	return nil
}
