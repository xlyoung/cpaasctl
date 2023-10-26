package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/config"
	"os"
	"path/filepath"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "cpctl",
	Short: "A brief description of your application",
	Long:  "",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(rootCmd)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "cpaasconfig", "", "config file (default is /opt/cpaas/config/cpaas.yaml)")

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

	// 使用确定的 configFile 初始化配置
	if err := config.InitConfig(configFile); err != nil {
		// 此处可以根据需要打印错误或执行其他错误处理
		fmt.Fprintf(os.Stderr, "Error initializing config: %s\n", err)
		os.Exit(1) // 或者其它您希望进行的错误处理
	}
}
