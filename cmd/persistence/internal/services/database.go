package services

import (
	"database/sql"
	"fmt"

	"github.com/FelipeMandelli/ProjetoInterativoV/cmd/persistence/internal/config"
	entitys "github.com/FelipeMandelli/ProjetoInterativoV/pkg/Entitys"
	_ "github.com/ibmdb/go_ibm_db"
	"github.com/spf13/viper"
)

const (
	newStudentProcedure    = "CALL insertStudent(?, ?, ?, ?, ?, ?)"
	newTeacherProcedure    = "CALL insertTeacher(?, ?)"
	newAttendanceProcedure = "CALL insertPresence(?, ?)"
)

// DATABASE=dbname;HOSTNAME=hostname;PORT=port;PROTOCOL=TCPIP;UID=username;PWD=passwd

func ConnectDatabase(provider *Provider) error {

	connString := fmt.Sprintf("HOSTNAME=%s;DATABASE=%s;PORT=%s;UID=%s;PWD=%s;security=ssl;",
		viper.GetString(config.DBHostKey),
		viper.GetString(config.DBNameKey),
		viper.GetString(config.DBPortKey),
		viper.GetString(config.DBUserKey),
		viper.GetString(config.DBPassKey))

	db, err := sql.Open("go_ibm_db", connString)
	if err != nil {
		return err
	}
	//defer db.Close()

	provider.DB = db
	provider.DbIsON = true

	return nil
}

func PersistAtendance(p *Provider, teacherTag, studentTag int) error {
	_, err := p.DB.Exec(newAttendanceProcedure, teacherTag, studentTag)
	if err != nil {
		return fmt.Errorf("error executing procedure: %w", err)
	}

	return nil
}

func PersistStudentRegistry(p *Provider, student entitys.Resgistry) error {
	_, err := p.DB.Exec(newStudentProcedure, student.Tag, student.Name, student.Document, student.Mail, student.Tel, student.Course)
	if err != nil {
		return fmt.Errorf("error executing procedure: %w", err)
	}

	return nil
}

func PersistTeacherRegistry(p *Provider, teacher entitys.Resgistry) error {
	_, err := p.DB.Exec(newTeacherProcedure, teacher.Tag, teacher.Name)
	if err != nil {
		return fmt.Errorf("error executing procedure: %w", err)
	}

	return nil
}
