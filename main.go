// 文件路径: main.go

package main

import (
	"github.com/spf13/cobra"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/cmdhandler" // 根据你的项目路径调整此行
	"log"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "cpctl",
		Short: "A CLI for controlling cp projects",
	}

	var configCmd = &cobra.Command{
		Use:   "config",
		Short: "Display the configuration",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmdhandler.DisplayConfig("config/cpaas.yaml")
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	rootCmd.AddCommand(configCmd)
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
