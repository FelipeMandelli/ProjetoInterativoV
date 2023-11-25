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
