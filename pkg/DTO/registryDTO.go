package dto

import "github.com/FelipeMandelli/ProjetoInterativoV/pkg/entities"

type RegistryDTO struct {
	SendingTime string
	Registry    entities.Registry
}

type SubjectRegistryDTO struct {
	SendingTime string
	Registry    entities.SubjectRegistry
}
