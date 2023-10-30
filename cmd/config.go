package cmd

import (
	"github.com/spf13/cobra"
	handler "gitlab.hycyg.com/paas-tools/cpaasctl/internal/handler"
	"log"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage app configurations",
}

var viewCmd = &cobra.Command{
	Use:   "view [appName]",
	Short: "Display the configuration for a specific app",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appName := args[0]
		err := handler.ViewConfig(appName)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var updateCmd = &cobra.Command{
	Use:   "update [appName] [configName] [configValue]",
	Short: "Update the configuration for a specific app",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		appName := args[0]
		configName := args[1]
		configValue := args[2]
		err := handler.UpdateConfig(cfgFile, appName, configName, configValue)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var exportCmd = &cobra.Command{
	Use:   "export ",
	Short: "export DockerCompose file",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		err := handler.ExportDockercomposeFile(cfgFile)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func SetupConfigCmd() {
	configCmd.AddCommand(viewCmd, updateCmd, exportCmd)
}
