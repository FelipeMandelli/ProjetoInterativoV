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
	postMehod = "POST"
)

func PackageSender(provider *Provider) {
	for {
		pack := <-provider.PackChan

		pack.SendingTime = time.Now().Local().Format(config.TimeFormater)

		provider.Log.Sugar().Infof("package to be sent: %+v", pack)

		status, err := sendPackToPersist(pack)
		if err != nil {
			provider.Log.Sugar().Error("error sending to Persist", err)
		}

		provider.Log.Sugar().Infof("status received: %s", status)
	}
}

func sendPackToPersist(pack dto.PackagerDTO) (string, error) {
	encodedPack, err := json.Marshal(pack)
	if err != nil {
		return "", fmt.Errorf("error encoding to json: %w", err)
	}

	req, err := http.NewRequest(postMehod, viper.GetString(config.PesistenceKey), bytes.NewBuffer(encodedPack))
	if err != nil {
		return "", fmt.Errorf("erro creating http request: %w", err)
	}

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error in request: %w", err)
	}

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return fmt.Errorf("error reading response: %w", err)
	// }

	resp.Body.Close()

	return resp.Status, nil
}
