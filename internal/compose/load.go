package compose

import (
	"fmt"
	"github.com/compose-spec/compose-go/cli"
	"github.com/compose-spec/compose-go/types"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/logger"
	"os"
)

func LoadAndInterpolateComposeFile(filePath string, environmentVars map[string]string) (*types.Project, error) {
	// 确保文件存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, err
	}

	// 准备环境变量
	var envs []string
	for key, value := range environmentVars {
		envs = append(envs, fmt.Sprintf("%s=%s", key, value))
	}

	// 创建 ProjectOptions
	options, err := cli.NewProjectOptions(
		[]string{filePath}, // 指定 docker-compose 文件的路径
		cli.WithOsEnv,      // 从操作系统加载环境变量
		cli.WithDotEnv,     // 加载 .env 文件
		cli.WithEnv(envs),  // 设置自定义环境变量
	)
	if err != nil {
		logger.Logger.Errorf("error creating project options: %v", err)
		return nil, err
	}

	// 从选项加载项目
	project, err := cli.ProjectFromOptions(options)
	if err != nil {
		logger.Logger.Errorf("error loading project from options: %v", err)
		return nil, err
	}

	return project, nil
}
