package services

import (
	"database/sql"
	"fmt"

	"github.com/FelipeMandelli/ProjetoInterativoV/cmd/persistence/internal/config"
	_ "github.com/ibmdb/go_ibm_db"
	"github.com/spf13/viper"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

// DATABASE=dbname;HOSTNAME=hostname;PORT=port;PROTOCOL=TCPIP;UID=username;PWD=passwd

func ConnectDatabase(provider *Provider) error {

	connString := fmt.Sprintf("DATABASE=%s;HOSTNAME=%s;PORT=%s;PROTOCOL=TCPIP;UID=%s;PWD=%s",
		// "{IBM DB2 ODBC DRIVER}",
		viper.GetString(config.DBNameKey),
		viper.GetString(config.DBHostKey),
		viper.GetString(config.DBPortKey),
		viper.GetString(config.DBUserKey),
		viper.GetString(config.DBPassKey))

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
	// 	viper.GetString(config.DBUserKey),
	// 	viper.GetString(config.DBPassKey),
	// 	viper.GetString(config.DBHostKey),
	// 	viper.GetString(config.DBPortKey),
	// 	viper.GetString(config.DBNameKey))

	provider.Log.Sugar().Infof("dsn string: %s", connString)

	// db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	// if err != nil {
	// 	return fmt.Errorf("failed to connect to database: %w", err)
	// }

	db, err := sql.Open("go_ibm_db", connString)
	if err != nil {
		return err
	}

	defer db.Close()

	// provider.DB = db

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
