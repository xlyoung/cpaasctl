package compose

import (
	"context"
	"fmt"
	"github.com/compose-spec/compose-go/loader"
	composeTypes "github.com/compose-spec/compose-go/types" // 添加别名以避免冲突
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/config"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/logger"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/utils"
	"io/ioutil"
)

func StartApp(appName string) error {
	// 加载配置文件引入变量
	cfg, err := config.LoadConfig("config/cpaas.yaml")
	if err != nil {
		return fmt.Errorf("error loading config: %w", err)
	}

	envVars, err := utils.SetEnvironmentVariables(cfg)
	if err != nil {
		return fmt.Errorf("error setting environment variables: %w", err)
	}

	// 读取docker-compose文件内容
	yamlFile, err := ioutil.ReadFile("./docker-compose.yml")
	if err != nil {
		logger.Logger.Errorf("Error reading docker-compose file: %v", err)
		return err
	}

	// 使用compose-go加载并解析docker-compose文件
	dict, err := loader.ParseYAML(yamlFile)
	if err != nil {
		logger.Logger.Errorf("Error parsing docker-compose file: %v", err)
		return err
	}

	// 注意这里使用了别名
	config, err := loader.Load(composeTypes.ConfigDetails{
		WorkingDir:  ".",
		ConfigFiles: []composeTypes.ConfigFile{{Config: dict}},
		Environment: envVars,
	})
	if err != nil {
		logger.Logger.Errorf("Error loading configuration: %v", err)
		return err
	}

	// 查找特定的服务配置
	var serviceConfig *composeTypes.ServiceConfig // 注意这里使用了别名
	for _, service := range config.Services {
		if service.Name == appName {
			serviceConfig = &service
			break
		}
	}

	if serviceConfig == nil {
		logger.Logger.Errorf("No service found with name: %s", appName)
		return fmt.Errorf("no service found with name: %s", appName)
	}

	// 创建Docker客户端
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		logger.Logger.Errorf("Cannot create Docker client: %v", err)
		return err
	}

	ctx := context.Background()

	// 拉取需要的镜像
	_, err = cli.ImagePull(ctx, serviceConfig.Image, types.ImagePullOptions{})
	if err != nil {
		logger.Logger.Errorf("Could not pull image %s: %v", serviceConfig.Image, err)
		return err
	}
	logger.Logger.Infof("Successfully pulled image %s", serviceConfig.Image)

	// 准备容器配置
	containerConfig := &container.Config{
		Image: serviceConfig.Image,
		Env:   envVarsToSlice(envVars), // 转换环境变量为字符串切片
		// ... 其他配置参数 ...
	}
	hostConfig := &container.HostConfig{
		// ... 主机配置参数 ...
	}

	// 创建并启动容器
	resp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, "")
	if err != nil {
		logger.Logger.Errorf("Cannot create container: %v", err)
		return err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		logger.Logger.Errorf("Cannot start container: %v", err)
		return err
	}

	logger.Logger.Infof("%s started successfully", appName)
	return nil
}

// envVarsToSlice converts a map of environment variables to a slice of strings.
func envVarsToSlice(envVars map[string]string) []string {
	result := make([]string, 0, len(envVars))
	for key, value := range envVars {
		result = append(result, fmt.Sprintf("%s=%s", key, value))
	}
	return result
}
