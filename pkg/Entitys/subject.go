package entitys

type Subject struct {
	Professor
	Classroom        string
	EnrolledStudents []Student
}
