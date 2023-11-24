package entities

import "gorm.io/gorm"

type Attendance struct {
	gorm.Model
	Date           string `gorm:"column:date"`
	Schedule       string `gorm:"column:schedule"`
	SubjectID      string `gorm:"column:subject_id"`
	AttendanceList string `gorm:"column:attendance_list"`
	AbsenceList    string `gorm:"column:absence_list"`
}
