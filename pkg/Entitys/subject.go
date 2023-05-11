package entitys

type Subject struct {
	Teacher
	Classroom        string
	EnrolledStudents []Student
}
