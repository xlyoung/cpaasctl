// 文件路径: internal/compose/load.go
package compose

import (
	"github.com/compose-spec/compose-go/loader"
	"github.com/compose-spec/compose-go/types"
	"os"
)

// ComposeConfig holds the docker compose configuration after loading it from file
type ComposeConfig struct {
	Config *types.Project
}

// LoadComposeFile reads a docker-compose.yml file and loads it into a ComposeConfig struct
func LoadComposeFile(filePath string, environmentVars map[string]string) (*ComposeConfig, error) {
	// Make sure the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, err
	}

	// Using the Compose Loader to load the configuration from file
	config, err := loader.Load(types.ConfigDetails{
		ConfigFiles: []types.ConfigFile{
			{
				Filename: filePath, // this is the path to your docker-compose.yml
				// If you have specific configuration content, you can specify it here
				// otherwise the loader will read the content from the provided file path
			},
		},
		// If there are any environment variables required for your config, ensure they are provided here
		Environment: environmentVars, // We are not passing any environment, but you can do it if required
	}, func(options *loader.Options) {
		//options.SkipSchemaCheck = true // This allows flexibility in the compose file version
	})

	if err != nil {
		return nil, err
	}

	return &ComposeConfig{Config: config}, nil
}

// Other functions that utilize ComposeConfig can be added here
