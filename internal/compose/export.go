package compose

import (
	"github.com/compose-spec/compose-go/types"
	"io/ioutil"
)

func ExportComposeFile(project *types.Project, newFilePath string) error {
	// 使用 MarshalYAML 方法将项目转换为 YAML 字符串
	yamlData, err := project.MarshalYAML()
	if err != nil {
		return err // 处理错误
	}

	// 将 YAML 数据写入新的 docker-compose 文件
	err = ioutil.WriteFile(newFilePath, []byte(yamlData), 0644)
	if err != nil {
		return err // 处理错误
	}

	return nil
}
