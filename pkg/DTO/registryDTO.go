package dto

import (
	"time"

	"github.com/FelipeMandelli/ProjetoInterativoV/pkg/entities"
)

type RegistryDTO struct {
	SendingTime time.Time
	Registry    entities.Registry
}

type SubjectRegistryDTO struct {
	SendingTime time.Time
	Registry    entities.SubjectRegistry
}
