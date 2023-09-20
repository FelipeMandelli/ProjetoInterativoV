package services

import (
	"fmt"

	"github.com/FelipeMandelli/ProjetoInterativoV/cmd/persistence/internal/config"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// const (
// 	newStudentProcedure    = "CALL insertStudent(?, ?, ?, ?, ?, ?)"
// 	newTeacherProcedure    = "CALL insertTeacher(?, ?)"
// 	newAttendanceProcedure = "CALL insertPresence(?, ?)"
// )

type Test struct {
	ID string `gorm:"coolumn:id"`
}

func ConnectDatabase(provider *Provider) error {

	db, err := gorm.Open(mysql.Open(createDBConnString(
		viper.GetInt(config.DBPortKey),
		viper.GetString(config.DBHostKey),
		viper.GetString(config.DBUserKey),
		viper.GetString(config.DBPassKey),
		viper.GetString(config.DBNameKey),
	)), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&Test{})
	if err != nil {
		return err
	}

	provider.DB = db
	provider.DbIsON = true

	return nil
}

// func PersistAtendance(p *Provider, teacherTag, studentTag string) error {
// 	_, err := p.DB.Exec(newAttendanceProcedure, teacherTag, studentTag)
// 	if err != nil {
// 		return fmt.Errorf("error executing procedure: %w", err)
// 	}

// 	return nil
// }

// func PersistStudentRegistry(p *Provider, student entitys.Resgistry) error {
// 	_, err := p.DB.Exec(newStudentProcedure, student.Tag, student.Name, student.Document, student.Mail, student.Tel, student.Course)
// 	if err != nil {
// 		return fmt.Errorf("error executing procedure: %w", err)
// 	}

// 	return nil
// }

// func PersistTeacherRegistry(p *Provider, teacher entitys.Resgistry) error {
// 	_, err := p.DB.Exec(newTeacherProcedure, teacher.Tag, teacher.Name)
// 	if err != nil {
// 		return fmt.Errorf("error executing procedure: %w", err)
// 	}

// 	return nil
// }

func createDBConnString(port int, host, username, password, name string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", username, password, host, port, name)
}
