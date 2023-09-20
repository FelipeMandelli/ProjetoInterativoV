package services

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var provider *Provider

type Provider struct {
	Log    *zap.Logger
	DB     *gorm.DB
	DbIsON bool
}

func GetProvider() *Provider {
	if provider != nil {
		return provider
	}

	provider = &Provider{
		DbIsON: false,
	}

	return provider
}
