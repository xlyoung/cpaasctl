// 文件路径: internal/cmdhandler/handler.go

package handler

import (
	"fmt"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/config" // 根据你的项目路径调整此行
)

func DisplayConfig(filePath string) error {
	cfg, err := config.LoadConfig(filePath)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", cfg)
	return nil
}
