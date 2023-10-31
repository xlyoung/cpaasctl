package compose

import (
	"bytes"
	"fmt"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/config"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/logger"
	"os"
	"os/exec"
	"path/filepath"
)

// DockerCompose 定义了管理 docker-compose 的操作
type DockerCompose interface {
	Up(target string) (string, error)
	Down(target string) (string, error)
	Status(target string) (string, error)
	Logs(target string) (string, error)
	Pull(target string) (string, error)
	Restart(target string) (string, error)
	// 根据需要添加其他方法
}

// DockerComposeImpl 是 DockerCompose 接口的具体实现
type DockerComposeImpl struct {
	Executable string            // docker-compose 可执行文件的路径
	ProjectDir string            // 项目的根目录
	EnvVars    map[string]string // 环境变量
}

// NewDockerComposeImpl 创建一个 DockerCompose 的新实例
func NewDockerComposeImpl(executablePath string, cfg *config.Config, envVars map[string]string) *DockerComposeImpl {
	projectDir := cfg.Cpaas.Base // 从配置中获取项目目录

	return &DockerComposeImpl{
		Executable: executablePath,
		ProjectDir: projectDir,
		EnvVars:    envVars,
	}
}

// checkDockerComposeFile 检查 docker-compose.yml 文件是否存在
func (dc *DockerComposeImpl) checkDockerComposeFile() error {
	dockerComposeFile := filepath.Join(dc.ProjectDir, "docker-compose.yml")
	if _, err := os.Stat(dockerComposeFile); os.IsNotExist(err) {
		return fmt.Errorf("docker-compose.yml 文件不存在: %s", dockerComposeFile)
	}
	return nil
}

// runDockerComposeCommand 是一个通用函数，用于执行 docker-compose 命令并捕获其输出
func (dc *DockerComposeImpl) runDockerComposeCommand(args []string) (string, error) {
	// 检查 docker-compose.yml 文件是否存在
	if err := dc.checkDockerComposeFile(); err != nil {
		return "", err
	}

	// 设置环境变量
	cmd := exec.Command(dc.Executable, args...)
	for key, value := range dc.EnvVars {
		cmd.Env = append(cmd.Env, key+"="+value)
	}

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("command failed: %v: %s", err, stderr.String())
	}
	logger.Logger.Infof("执行命令: %s %s\n", dc.Executable, args)
	if stdout.Len() > 0 {
		logger.Logger.Infof("命令stdout输出:\n%s\n", stdout.String())
	}
	if stderr.Len() > 0 {
		logger.Logger.Infof("命令stderr输出:\n%s\n", stderr.String())
	}
	return stdout.String(), nil
}

// Up 启动服务。如果指定了 target，则只启动该服务；否则，启动所有服务。
func (dc *DockerComposeImpl) Up(target string) (string, error) {
	args := []string{"-f", filepath.Join(dc.ProjectDir, "docker-compose.yml"), "up", "-d"}
	if target != "" {
		args = append(args, target) // 如果指定了 target，则添加到命令参数中
	}

	return dc.runDockerComposeCommand(args)
}

// Down 停止服务
func (dc *DockerComposeImpl) Down(target string) (string, error) {
	args := []string{"-f", filepath.Join(dc.ProjectDir, "docker-compose.yml"), "down"}
	if target != "" {
		args = append(args, target)
	}
	return dc.runDockerComposeCommand(args)
}

// Status 检查服务的状态
func (dc *DockerComposeImpl) Status(target string) (string, error) {
	args := []string{"-f", filepath.Join(dc.ProjectDir, "docker-compose.yml"), "ps"}
	if target != "" {
		args = append(args, target)
	}
	return dc.runDockerComposeCommand(args)
}

// Logs 获取服务的日志
func (dc *DockerComposeImpl) Logs(target string) (string, error) {
	args := []string{"-f", filepath.Join(dc.ProjectDir, "docker-compose.yml"), "logs"}
	if target != "" {
		args = append(args, target)
	}
	return dc.runDockerComposeCommand(args)
}

// Pull 拉取服务的最新镜像
func (dc *DockerComposeImpl) Pull(target string) (string, error) {
	args := []string{"-f", filepath.Join(dc.ProjectDir, "docker-compose.yml"), "pull"}
	if target != "" {
		args = append(args, target)
	}
	return dc.runDockerComposeCommand(args)
}

// Restart 重启服务
func (dc *DockerComposeImpl) Restart(target string) (string, error) {
	// 停止服务
	if _, err := dc.Down(target); err != nil {
		return "", err
	}

	// 启动服务
	return dc.Up(target)
}
