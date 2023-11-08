package cmd

import (
	"github.com/spf13/cobra"
	handler "gitlab.hycyg.com/paas-tools/cpaasctl/internal/handler"
	"log"
)

var newappHandler *handler.AppHandler

func getAppHandler() (*handler.AppHandler, error) {
	if newappHandler == nil {
		var err error
		newappHandler, err = handler.NewAppHandler()
		if err != nil {
			return nil, err
		}
	}
	return newappHandler, nil
}

func SetupAppCmd() {
	startCmd := &cobra.Command{
		Use:   "start [appName]",
		Short: "Start an application",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			newappHandler, err := getAppHandler()
			if err != nil {
				log.Fatal(err)
			}
			appName := ""
			if len(args) > 0 {
				appName = args[0]
			}
			err = newappHandler.StartApp(appName)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	stopCmd := &cobra.Command{
		Use:   "stop [appName]",
		Short: "Stop an application",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			newappHandler, err := getAppHandler()
			if err != nil {
				log.Fatal(err)
			}
			appName := ""
			if len(args) > 0 {
				appName = args[0]
			}
			err = newappHandler.StopApp(appName)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	restartCmd := &cobra.Command{
		Use:   "restart [appName]",
		Short: "Restart an application",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			newappHandler, err := getAppHandler()
			if err != nil {
				log.Fatal(err)
			}
			appName := ""
			if len(args) > 0 {
				appName = args[0]
			}
			err = newappHandler.RestartApp(appName)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	statusCmd := &cobra.Command{
		Use:   "status [appName]",
		Short: "Get status of an application",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			newappHandler, err := getAppHandler()
			if err != nil {
				log.Fatal(err)
			}
			appName := ""
			if len(args) > 0 {
				appName = args[0]
			}
			err = newappHandler.GetAppStatus(appName)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	pullCmd := &cobra.Command{
		Use:   "pull [appName]",
		Short: "pull images of an application",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			newappHandler, err := getAppHandler()
			if err != nil {
				log.Fatal(err)
			}
			appName := ""
			if len(args) > 0 {
				appName = args[0]
			}
			err = newappHandler.PullApp(appName)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	logCmd := &cobra.Command{
		Use:   "log [appName]",
		Short: "log of an application",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			newappHandler, err := getAppHandler()
			if err != nil {
				log.Fatal(err)
			}
			appName := ""
			if len(args) > 0 {
				appName = args[0]
			}
			err = newappHandler.LogApp(appName)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	rootCmd.AddCommand(startCmd, stopCmd, restartCmd, statusCmd, pullCmd, logCmd) // 将生成的子命令添加到 appCmd
}
