package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	basePath     = "config"
	AddressKey   = "ADDRESS"
	TimeFormater = "02-01-2006 15:04:05"
)

func SetupConfigurations() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(basePath)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading configs: %w", err)
	}

	return nil
}
