package handler

import (
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/config"
	logger "gitlab.hycyg.com/paas-tools/cpaasctl/internal/logger"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/pkg/compose"
	composeImpl "gitlab.hycyg.com/paas-tools/cpaasctl/internal/pkg/compose"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/pkg/utils"
)

// AppHandler 是应用处理程序
type AppHandler struct {
	cfg           *config.Config
	envVars       map[string]string
	dockerCompose compose.DockerCompose
}

// NewAppHandler 创建一个新的 AppHandler 实例
func NewAppHandler() (*AppHandler, error) {
	cfg := config.GetConfig()
	envVars, err := utils.SetEnvironmentVariables(cfg)
	if err != nil {
		return nil, err
	}

	dockerComposeBinPath, err := utils.FindDockerCompose()
	if err != nil {
		return nil, err
	}

	dockerCompose := composeImpl.NewDockerComposeImpl(dockerComposeBinPath, cfg, envVars)

	return &AppHandler{
		cfg:           cfg,
		envVars:       envVars,
		dockerCompose: dockerCompose,
	}, nil
}

// StartApp 启动应用
func (ah *AppHandler) StartApp(appName string) error {
	logger.Logger.Infof("Starting %s...\n", appName)

	// 使用接口启动服务
	if _, err := ah.dockerCompose.Up(appName); err != nil {
		return err
	}

	return nil
}

// StopApp 停止应用
func (ah *AppHandler) StopApp(appName string) error {
	logger.Logger.Infof("Stopping %s...\n", appName)

	// 使用接口启动服务
	if _, err := ah.dockerCompose.Down(appName); err != nil {
		return err
	}
	// 实现应用的停止逻辑，例如，使用 Docker SDK 停止容器或发送请求到一个 API

	return nil
}

// RestartApp 重启应用
func (ah *AppHandler) RestartApp(appName string) error {
	logger.Logger.Infof("Restarting %s...\n", appName)

	// 使用接口启动服务
	if _, err := ah.dockerCompose.Restart(appName); err != nil {
		return err
	}

	return nil
}

// GetAppStatus 获取应用状态
func (ah *AppHandler) GetAppStatus(appName string) error {
	logger.Logger.Infof("Getting status of %s...\n", appName)

	// 使用接口启动服务
	if _, err := ah.dockerCompose.Status(appName); err != nil {
		return err
	}
	// 实现获取应用状态的逻辑，例如，查询 Docker 容器的状态或发送请求到一个 API

	return nil
}

// PullApp 拉取镜像
func (ah *AppHandler) PullApp(appName string) error {
	logger.Logger.Infof("Pulling image of %s...\n", appName)

	// 使用接口启动服务
	if _, err := ah.dockerCompose.Pull(appName); err != nil {
		return err
	}
	// 实现获取应用状态的逻辑，例如，查询 Docker 容器的状态或发送请求到一个 API

	return nil
}

// LogApp 显示日志
func (ah *AppHandler) LogApp(appName string) error {
	logger.Logger.Infof("Logging of %s...\n", appName)

	// 使用接口启动服务
	if _, err := ah.dockerCompose.Logs(appName); err != nil {
		return err
	}

	return nil
}
