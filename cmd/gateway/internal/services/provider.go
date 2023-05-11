package services

import "go.uber.org/zap"

var provider *Provider

type Provider struct {
	Log *zap.Logger
}

func GetProvider() *Provider {
	if provider != nil {
		return provider
	}

	provider = &Provider{}

	return provider
}
