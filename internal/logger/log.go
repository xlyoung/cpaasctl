// logger.go
package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

// Logger 是我们应用程序中使用的标准日志记录器
var Logger = logrus.New()

func init() {
	// 设置日志级别，生产环境中可能是 Info 或 Warn，开发环境中可以是 Debug
	Logger.SetLevel(logrus.InfoLevel) // 这可以通过环境变量或配置来控制

	// 设置输出，它可以是任何 io.Writer，这里我们用标准输出
	Logger.SetOutput(os.Stdout)

	// 设置日志格式，我们这里用的是 JSON，但你也可以使用 logrus.TextFormatter
	Logger.SetFormatter(&logrus.TextFormatter{})
}
