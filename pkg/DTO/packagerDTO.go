package dto

import "time"

type PackagerDTO struct {
	SendingTime   time.Time
	TeacherID     string
	AttendanceIDs []string
}
