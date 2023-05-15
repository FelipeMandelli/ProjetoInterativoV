package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	basePath     = "config"
	AddressKey   = "ADDRESS"
	DBUserKey    = "DB_USER"
	DBPassKey    = "DB_PASS"
	DBHostKey    = "DB_HOST"
	DBPortKey    = "DB_PORT"
	DBNameKey    = "DB_NAME"
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
