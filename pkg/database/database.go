package database

import "time"

type Course struct {
	ID   int //pk
	Name string
}

type Student struct {
	ID       int //pk
	Tag      int
	Name     string
	Document string
	Mail     string
	Phone    string
	Course   int //fk
}

type Teacher struct {
	ID   int //pk
	Tag  int
	Name string
}

type Class struct {
	ID        int //pk
	Classroom string
	Curso     string //fk
	Subject   string //fk
	TeacherID int    //fk
	DateTime  time.Time
}

type Subject struct {
	ID        int //pk
	Name      string
	TeacherID int //fk
	Course    int //fk
}

type Attendance struct {
	ID         int //pk
	TeacherTag int
	StudentTag int
	TeacherID  int //fk
	StudentID  int //fk
	ClassID    int //fk
	SubjectID  int //fk
	CourseID   int //fk
	DateTime   time.Time
}
