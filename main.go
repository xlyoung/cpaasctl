package main

import (
	"github.com/spf13/cobra"
	handler "gitlab.hycyg.com/paas-tools/cpaasctl/internal/handler"
	"log"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "cpctl",
		Short: "A CLI for controlling cp projects",
	}

	// configCmd represents the config command
	// 创建 'config' 子命令
	var configCmd = &cobra.Command{
		Use:   "config",
		Short: "Manage app configurations",
		// 这里不再需要 'Run'，因为 'config' 现在有自己的子命令
	}

	// 创建 'view' 子命令
	var viewCmd = &cobra.Command{
		Use:   "view [appName]",
		Short: "Display the configuration for a specific app",
		Args:  cobra.ExactArgs(1), // 需要一个参数: appName
		Run: func(cmd *cobra.Command, args []string) {
			appName := args[0]
			err := handler.ViewConfig("config/cpaas.yaml", appName)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	// 创建 'update' 子命令
	var updateCmd = &cobra.Command{
		Use:   "update [appName] [configName] [configValue]",
		Short: "Update the configuration for a specific app",
		Args:  cobra.ExactArgs(3), // 需要三个参数: appName, configName, 和 configValue
		Run: func(cmd *cobra.Command, args []string) {
			appName := args[0]
			configName := args[1]
			configValue := args[2]
			// 假设您有一个 'UpdateConfig' 函数来处理更新
			err := handler.UpdateConfig("config/cpaas.yaml", appName, configName, configValue)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	// 创建app命令，这将作为其他子命令的父命令
	var appCmd = &cobra.Command{
		Use:   "app",
		Short: "Manage applications",
	}

	// 创建start子命令
	var startCmd = &cobra.Command{
		Use:   "start [appName]",
		Short: "Start an application",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			appName := args[0]
			err := handler.StartApp(appName)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	// 创建stop子命令
	var stopCmd = &cobra.Command{
		Use:   "stop [appName]",
		Short: "Stop an application",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			appName := args[0]
			err := handler.StopApp(appName)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	// 创建restart子命令
	var restartCmd = &cobra.Command{
		Use:   "restart [appName]",
		Short: "Restart an application",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			appName := args[0]
			err := handler.RestartApp(appName)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	// 创建status子命令
	var statusCmd = &cobra.Command{
		Use:   "status [appName]",
		Short: "Get the status of an application",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			appName := args[0]
			err := handler.GetAppStatus(appName)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	// 将start, stop, restart, status子命令添加到app命令
	appCmd.AddCommand(startCmd, stopCmd, restartCmd, statusCmd)

	// 将 'view' 和 'update' 子命令添加到 'config' 命令
	configCmd.AddCommand(viewCmd)
	configCmd.AddCommand(updateCmd)

	// 将 'config' 命令添加到根命令
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(appCmd)

	// 执行根命令
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
