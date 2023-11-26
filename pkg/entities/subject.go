package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Subject struct {
	gorm.Model
	Name             string `gorm:"column:name"`
	Semester         string    `gorm:"column:reference_semester"`
	Year             string `gorm:"column:reference_year"`
	ProfessorName    string `gorm:"column:professor_name"`
	ProfessorID      string `gorm:"primary_key;column:professor_id"`
	StudentsEnrolled string `gorm:"column:students_enrolled_ids"`
	WeekDay          string    `gorm:"primary_key;column:week_day"`
	Schedule         string    `gorm:"primary_key;column:schedule"`
}

func (s *Subject) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == 0 {
		s.ID = uint(uuid.New().ID())
	}

	return
}
