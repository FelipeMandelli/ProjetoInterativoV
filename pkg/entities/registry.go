package entities

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
