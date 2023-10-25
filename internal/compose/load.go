package compose

import (
	"fmt"
	"github.com/compose-spec/compose-go/interpolation"
	"github.com/compose-spec/compose-go/loader"
	"github.com/compose-spec/compose-go/types"
	"os"
)

func LoadAndInterpolateComposeFile(filePath string, environmentVars map[string]string) (*types.Project, error) {
	// 确保文件存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, err
	}

	// 使用Compose Loader从文件加载配置
	configDetails := types.ConfigDetails{
		ConfigFiles: []types.ConfigFile{
			{
				Filename: filePath,
			},
		},
		Environment: environmentVars,
	}

	config, err := loader.Load(configDetails)
	if err != nil {
		return nil, err
	}

	lookupEnv := func(key string) (string, bool) {
		value, exists := environmentVars[key]
		return value, exists
	}

	interpOptions := interpolation.Options{
		LookupValue: lookupEnv,
	}

	interpolatedConfig, err := interpolation.Interpolate(config, interpOptions)
	if err != nil {
		return nil, fmt.Errorf("error interpolating variables in docker-compose file: %v", err)
	}

	return loader.Load(types.ConfigDetails{
		WorkingDir:  ".",
		ConfigFiles: []types.ConfigFile{{Config: interpolatedConfig}},
		Environment: environmentVars,
	})
}
