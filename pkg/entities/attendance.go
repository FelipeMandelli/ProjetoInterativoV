package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Attendance struct {
	gorm.Model
	Date           string `gorm:"primary_key;column:date"`
	Schedule       string `gorm:"primary_key;column:schedule"`
	ProfessorID    string `gorm:"primary_key;column:professor_id"`
	SubjectID      string `gorm:"column:subject_id"`
	AttendanceList string `gorm:"column:attendance_list"`
	AbsenceList    string `gorm:"column:absence_list"`
}

func (s *Attendance) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == 0 {
		s.ID = uint(uuid.New().ID())
	}

	return
}
