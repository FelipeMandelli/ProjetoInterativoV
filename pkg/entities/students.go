package entities

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	IDBiometry      string `gorm:"primary_key;column:id_biometry"`
	Name            string `gorm:"column:name"`
	Email           string `gorm:"column:email"`
	Course          string `gorm:"column:course"`
	CellphoneNumber string `gorm:"column:cellphone_number"`
}
