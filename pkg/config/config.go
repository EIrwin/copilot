package config

import (
	"os"

	"github.com/spf13/viper"
)

const (
	configPathEnv = "CONFIG_PATH"
	portEnv       = "PORT"
	slackTokenEnv = "SLACK_TOKEN"
)

func Init(configPath string) error {

	viper.Set(configPathEnv, parseConfigPath(configPath))
	viper.Set(portEnv, os.Getenv(portEnv))
	viper.Set(slackTokenEnv, os.Getenv(slackTokenEnv))

	return validate()
}

func ConfigPath() string {
	return viper.GetString(configPathEnv)
}

func Port() string {
	return viper.GetString(portEnv)
}

func SlackToken() string {
	return viper.GetString(slackTokenEnv)
}

func parseConfigPath(path string) string {
	if path != "" {
		return path
	}
	return os.Getenv(configPathEnv)
}

func validate() error {
	if Port() == "" {
		return NewErrInvalidConfig(portEnv)
	}

	if SlackToken() == "" {
		return NewErrInvalidConfig(slackTokenEnv)
	}

	return nil
}
