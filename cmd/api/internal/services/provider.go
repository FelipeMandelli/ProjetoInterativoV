package services

import (
	domain "github.com/FelipeMandelli/ProjetoInterativoV/cmd/api/internal/domain/rest"
	dto "github.com/FelipeMandelli/ProjetoInterativoV/pkg/DTO"
	"go.uber.org/zap"
)

var provider *Provider

type Provider struct {
	Log             *zap.Logger
	PackChan        chan dto.PackagerDTO
	RequestBodyChan chan domain.AttendanceRequest
	RegChan         chan dto.RegistryDTO
}

func GetProvider() *Provider {
	if provider != nil {
		return provider
	}

	provider = &Provider{
		PackChan:        make(chan dto.PackagerDTO),
		RequestBodyChan: make(chan domain.AttendanceRequest),
		RegChan:         make(chan dto.RegistryDTO),
	}

	return provider
}
