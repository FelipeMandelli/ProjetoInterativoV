package entities

import "errors"

type Registry struct {
	IDBiometry      string `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"mail"`
	Course          string `json:"course"`
	CellphoneNumber string `json:"cellphone"`
	Role            string `json:"role"`
}

type Role string
type CoursesAccepted string

const (
	StudentRole   Role = "S"
	ProfessorRole Role = "P"
)

const (
	ComputerEngineering   CoursesAccepted = "Engenharia da Computação"
	IndustrialEngineering CoursesAccepted = "Engenharia de Produção"
)

func (r *Registry) IsValidRole() bool {
	if r.Role != string(StudentRole) && r.Role != string(ProfessorRole) {
		return false
	}

	return true
}

func (r *Registry) IsValidCourse() bool {
	if r.Course != string(ComputerEngineering) && r.Course != string(IndustrialEngineering) {
		return false
	}

	return true
}

func (r *Registry) ToStudent() (*Student, error) {
	if r.Role != string(StudentRole) {
		return nil, errors.New("cannot save a non student into student")
	}

	return &Student{
		IDBiometry:      r.IDBiometry,
		Name:            r.Name,
		Email:           r.Email,
		Course:          r.Course,
		CellphoneNumber: r.CellphoneNumber,
	}, nil
}

func (r *Registry) ToProfessor() (*Professor, error) {
	if r.Role != string(ProfessorRole) {
		return nil, errors.New("cannot save a non student into student")
	}

	return &Professor{
		IDBiometry:      r.IDBiometry,
		Name:            r.Name,
		Email:           r.Email,
		CellphoneNumber: r.CellphoneNumber,
	}, nil
}
