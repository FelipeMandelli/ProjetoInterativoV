package services

import (
	"fmt"

	"github.com/FelipeMandelli/ProjetoInterativoV/cmd/persistence/internal/config"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase(provider *Provider) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		viper.GetString(config.DBUserKey),
		viper.GetString(config.DBPassKey),
		viper.GetString(config.DBHostKey),
		viper.GetString(config.DBPortKey),
		viper.GetString(config.DBNameKey))

	provider.Log.Sugar().Infof("dsn string: %s", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	provider.DB = db

	return nil
}
