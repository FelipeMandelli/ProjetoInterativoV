package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	basePath   = "cmd/api/config"
	AddressKey = "ADDRESS"
)

func SetupConfigurations() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(basePath)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading configs: %w", err)
	}

	return nil
}
