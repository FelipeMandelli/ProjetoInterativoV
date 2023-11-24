package entities

import "gorm.io/gorm"

type Professor struct {
	gorm.Model
	IDBiometry      string `gorm:"primary_key;column:id_biometry"`
	Name            string `gorm:"column:name"`
	Email           string `gorm:"column:email"`
	CellphoneNumber string `gorm:"column:cellphone_number"`
}
