package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Professor struct {
	gorm.Model
	IDBiometry      string `gorm:"primary_key;column:id_biometry"`
	Name            string `gorm:"column:name"`
	Email           string `gorm:"column:email"`
	CellphoneNumber string `gorm:"column:cellphone_number"`
}

func (s *Professor) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == 0 {
		s.ID = uint(uuid.New().ID())
	}

	return
}
