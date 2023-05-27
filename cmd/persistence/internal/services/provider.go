package services

import (
	"database/sql"

	"go.uber.org/zap"
)

var provider *Provider

type Provider struct {
	Log    *zap.Logger
	DB     *sql.DB
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
