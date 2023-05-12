package services

import (
	"time"

	"github.com/FelipeMandelli/ProjetoInterativoV/cmd/gateway/internal/config"
)

func PackageSender(provider *Provider) {
	for {
		pack := <-provider.PackChan

		pack.SendingTime = time.Now().Local().Format(config.TimeFormater)

		provider.Log.Sugar().Infof("package to be sent: %+v", pack)
	}
}
