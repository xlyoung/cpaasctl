package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	handler "gitlab.hycyg.com/paas-tools/cpaasctl/internal/handler"
	"log"
)

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Manage applications",
}

func SetupAppCmd() {
	newappHandler, err := handler.NewAppHandler() // 初始化 newappHandler
	if err != nil {
		log.Fatal(err)
	}

	appCommands := map[string]func(appName string) error{
		"start":   newappHandler.StartApp,
		"stop":    newappHandler.StopApp,
		"restart": newappHandler.RestartApp,
		"status":  newappHandler.GetAppStatus,
	}

	for cmdName, cmdFunc := range appCommands {
		cmd := &cobra.Command{
			Use:   fmt.Sprintf("%s [appName]", cmdName),
			Short: fmt.Sprintf("%s an application", cmdName),
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				appName := args[0]
				err := cmdFunc(appName)
				if err != nil {
					log.Fatal(err)
				}
			},
		}
		appCmd.AddCommand(cmd) // 将生成的子命令添加到 appCmd
	}
}
