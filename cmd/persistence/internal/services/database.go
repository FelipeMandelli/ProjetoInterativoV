package services

import (
	"fmt"

	"github.com/FelipeMandelli/ProjetoInterativoV/cmd/persistence/internal/config"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase(p *Provider) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		viper.GetString(config.DBUserKey),
		viper.GetString(config.DBPassKey),
		viper.GetString(config.DBHostKey),
		viper.GetString(config.DBPortKey),
		viper.GetString(config.DBNameKey))

	p.Log.Sugar().Infof("dsn string: %s", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	p.DB = db
	p.DbIsON = true

	return nil
}

func PersistAtendance(p *Provider, param1 string, param2 int) error {

	// result := p.DB.Raw("CALL YourProcedureName(?, ?)", param1, param2)
	// if result.Error != nil {
	// 	return fmt.Errorf("error calling procedure: %w", result.Error)
	// }

	return nil
}

func PersistRegistry(p *Provider) error {
	return nil
}
