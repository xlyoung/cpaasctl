package utils

import (
	"errors"
	"fmt"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/config"
	logger "gitlab.hycyg.com/paas-tools/cpaasctl/internal/logger"
	"os"
	"os/exec"
	"strings"
)

// SetEnvironmentVariables now returns a map of the environment variables
func SetEnvironmentVariables(cfg *config.Config) (map[string]string, error) {
	envVars := make(map[string]string)

	// Set CPAAS_BASE
	baseKey := ConvertToEnvVarName("cpaas.base")
	baseValue := cfg.Cpaas.Base
	envVars[baseKey] = baseValue
	logger.Logger.Debugf("Set %s = %s\n", baseKey, baseValue)

	// Set CPAAS_REGISTRY_URL
	registryURLKey := ConvertToEnvVarName("cpaas.registry.url")
	registryURLValue := cfg.Cpaas.Registry.URL
	envVars[registryURLKey] = registryURLValue
	logger.Logger.Debugf("Set %s = %s\n", registryURLKey, registryURLValue)

	// For the App section, you might want to set a version environment variable for each application
	for appName, appConfig := range cfg.App {
		// Create a key like APP_STORAGE_VERSION
		versionKey := ConvertToEnvVarName(appName + ".version")
		versionValue := appConfig.Version
		envVars[versionKey] = versionValue
		logger.Logger.Debugf("Set %s = %s\n", versionKey, versionValue)
	}

	return envVars, nil
}

// ConvertToEnvVarName takes a string and converts it into a format suitable for an environment variable.
// This is generally UPPER_CASE format.
func ConvertToEnvVarName(key string) string {
	upperKey := strings.ToUpper(key)
	envKey := strings.ReplaceAll(upperKey, "-", "_")
	return strings.ReplaceAll(envKey, ".", "_")
}

func MapToStringSlice(envVars map[string]string) []string {
	var stringSlice []string
	for key, value := range envVars {
		stringSlice = append(stringSlice, fmt.Sprintf("%s=%s", key, value))
	}
	return stringSlice
}

func FindDockerCompose() (string, error) {
	// Check for the DOCKER_COMPOSE_PATH environment variable, which the user can set to specify a custom path.
	if path := os.Getenv("DOCKER_COMPOSE_PATH"); path != "" {
		// Verify that the file exists at the specified path.
		if _, err := os.Stat(path); err == nil {
			return path, nil
		} else if errors.Is(err, os.ErrNotExist) {
			return "", errors.New("docker-compose not found at the specified DOCKER_COMPOSE_PATH")
		} else {
			return "", err // Other error (e.g., permission denied)
		}
	}

	// If the environment variable is not set, try to locate it using the system's PATH.
	path, err := exec.LookPath("docker-compose")
	if err != nil {
		return "", errors.New("docker-compose not found in the system's PATH")
	}
	return path, nil
}
