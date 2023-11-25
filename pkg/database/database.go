package database

import (
	"fmt"

	"github.com/FelipeMandelli/ProjetoInterativoV/pkg/entities"
	"gorm.io/gorm"
)

func SaveStudent(db *gorm.DB, student *entities.Student) error {
	err := db.Save(student).Error
	if err != nil {
		return fmt.Errorf("error saving student: [%w]", err)
	}

	return err
}

func SaveProfessor(db *gorm.DB, professor *entities.Professor) error {
	err := db.Save(professor).Error
	if err != nil {
		return fmt.Errorf("error saving professor: [%w]", err)
	}

	return err
}

func SaveSubject(db *gorm.DB, subject *entities.Subject) error {
	err := db.Save(subject).Error
	if err != nil {
		return fmt.Errorf("error saving subject: [%w]", err)
	}

	return err
}

func SaveAttendance(db *gorm.DB, attendance *entities.Attendance) error {
	err := db.Save(attendance).Error
	if err != nil {
		return fmt.Errorf("error saving attendance: [%w]", err)
	}

	return err
}
