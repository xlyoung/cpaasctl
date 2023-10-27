package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/config"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/logger"
	"os"
	"path/filepath"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "A brief description of your application",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run is being called")
	},
}
var cfgFile string

func Execute() {
	// 设置标记

	setupFlags()
	setupCommands()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "执行错误: %s\n", err)
		os.Exit(1)
	}
}

func setupFlags() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "cpaasconfig", "", "config file (default is /opt/cpaas/config/cpaas.yaml)")
}

func setupCommands() {

	SetupConfigCmd()
	rootCmd.AddCommand(configCmd)
	SetupAppCmd()
	rootCmd.AddCommand(appCmd)

}

func initConfig() {
	// 预先定义好 configFile 变量，我们稍后会根据不同的条件来为其赋值
	var configFile string

	if cfgFile != "" {
		// 使用命令行参数指定的配置文件。
		configFile = cfgFile
	} else if cpaasEnvConfig := os.Getenv("CPAASCONFIG"); cpaasEnvConfig != "" {
		// 使用环境变量指定的配置文件路径。
		configFile = cpaasEnvConfig
	} else {
		// 设置默认配置文件路径。
		// 首先尝试相对于可执行文件的路径。
		exeDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		fmt.Println("exeDir", exeDir)
		if err != nil {
			cobra.CheckErr(err)
		}
		relativePath := filepath.Join(exeDir, "config", "cpaas.yaml")

		if _, err := os.Stat(relativePath); !os.IsNotExist(err) {
			// 如果相对路径下的文件存在，则使用该路径。
			configFile = relativePath
		} else {
			// 如果相对路径没有找到文件，则使用默认路径。
			configFile = "/opt/cpaas/config/cpaas.yaml"
		}
	}
	logger.Logger.Infof("Using config file: %s\n", configFile)
	// 使用确定的 configFile 初始化配置
	if err := config.InitConfig(configFile); err != nil {
		// 此处可以根据需要打印错误或执行其他错误处理
		fmt.Fprintf(os.Stderr, "Error initializing config: %s\n", err)
		os.Exit(1) //
	}
}
