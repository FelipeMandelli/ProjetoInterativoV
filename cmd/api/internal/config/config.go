package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

const (
	basePath           = "config"
	AddressKey         = "ADDRESS"
	PesistenceKey      = "PERSISTENCE_ADDRES"
	AttendancePathKey  = "ATTENDANCE_PATH"
	RegistryPathKey    = "REGISTRY_PATH"
	SubRegistryPathKey = "SUB_REGISTRY_PATH"
	RetryKey           = "PERSISTENCE_RETRY"
)

var (
	SPTimeZone *time.Location
)

func SetupConfigurations() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(basePath)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading configs: %w", err)
	}

	SPTimeZone, _ = time.LoadLocation("America/Sao_Paulo")

	return nil
}
