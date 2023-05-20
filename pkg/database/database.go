package database

type Course struct {
  ID int //pk
  Name String
}

type Student struct {
  ID int //pk
  Tag int
  Name String
  Document string
  Mail String
  Phone string
  Course int //fk
}

type Teacher struct {
  ID int //pk
  Tag int
  Name String
}

type Class struct {
  ID int //pk
  Classroom String
  Curso string //fk
  Subject string //fk
  TeacherID int //fk
  DateTime time.datetime
}

type Subject struct {
  ID int //pk
  Name string
  TeacherID int //fk
  Course int //fk
}

type Attendance struct {
  ID int //pk
  TeacherTag int
  StudentTag int
  TeacherID int //fk
  StudentID int //fk
  ClassID int //fk
  SubjectID int //fk
  CourseID int //fk
  DateTime time.datetime
}
