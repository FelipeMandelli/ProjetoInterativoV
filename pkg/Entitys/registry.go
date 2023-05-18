package entitys

type Resgistry struct {
	Name   string
	Mail   string
	Role   Role
	Course Course
}

type Role string

type Course string

const (
	TeacherRole      Role   = "Professor"
	StudentRole      Role   = "Aluno"
	ComputingCourse  Course = "Engenharia da Computação"
	ProductionCourse Course = "Engenharia de Produção"
)

func (r *Resgistry) IsValidRole() bool {
	return r.Role == TeacherRole || r.Role == StudentRole
}

func (r *Resgistry) IsValidCourse() bool {
	return r.Course == ComputingCourse || r.Course == ProductionCourse
}
