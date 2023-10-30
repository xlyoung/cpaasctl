package compose

import (
	"github.com/compose-spec/compose-go/types"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/logger"
	"os"
)

func ExportComposeFile(project *types.Project, newFilePath string) error {
	// 使用 MarshalYAML 方法将项目转换为 YAML 字符串
	yamlData, err := project.MarshalYAML()
	if err != nil {
		return err // 处理错误
	}

	// 创建新的文件并打开以供写入
	file, err := os.Create(newFilePath)
	if err != nil {
		return err // 处理错误
	}
	// 使用 defer 延迟关闭文件，以确保无论如何都会关闭文件
	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = cerr // 如果关闭文件时出现错误，将其保存在 err 中
		}
	}()

	// 将 YAML 数据写入新的 docker-compose 文件
	_, err = file.Write(yamlData)
	if err != nil {
		return err // 处理错误
	}

	// 刷新缓冲并确保写入到磁盘
	if err := file.Sync(); err != nil {
		return err // 处理错误
	}

	logger.Logger.Infof("Exported docker-compose file to %s\n", newFilePath)
	return nil
}
