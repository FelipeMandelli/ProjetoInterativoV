package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/FelipeMandelli/ProjetoInterativoV/cmd/api/internal/config"
	dto "github.com/FelipeMandelli/ProjetoInterativoV/pkg/DTO"
	"github.com/spf13/viper"
)

const (
	postMehod     = "POST"
	retryInterval = 2
)

func PackageSender(provider *Provider) {
	for {
		pack := <-provider.PackChan

		pack.SendingTime = time.Now().Local().Format(config.TimeFormater)

		provider.Log.Sugar().Infof("package to be sent: %+v", pack)

		status, err := sendPackToPersist(provider, pack)
		if err != nil {
			provider.Log.Sugar().Error("error sending to Persist: ", err)
		}

		provider.Log.Sugar().Infof("status received: %s", status)
	}
}

func sendPackToPersist(p *Provider, pack dto.PackagerDTO) (string, error) {
	encodedPack, err := json.Marshal(pack)
	if err != nil {
		return "", fmt.Errorf("error encoding to json: %w", err)
	}

	req, err := http.NewRequest(postMehod, viper.GetString(config.PesistenceKey), bytes.NewBuffer(encodedPack))
	if err != nil {
		return "", fmt.Errorf("erro creating http request: %w", err)
	}

	client := http.DefaultClient

	var resp *http.Response

	for i := 1; i <= viper.GetInt(config.RetryKey); i++ {
		p.Log.Sugar().Infof("trying to send request %d", i)

		resp, err = client.Do(req)
		if err == nil && resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
			return resp.Status, nil
		}

		time.Sleep(time.Second * retryInterval)
	}
	return "", fmt.Errorf("error while making http request: %w", err)
}

func RegistrySender(provider *Provider) {

}
