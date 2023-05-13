package services

import (
	dto "github.com/FelipeMandelli/ProjetoInterativoV/pkg/DTO"
)

// Packager packages attendance info in DTO
func Packager(provider *Provider) {
	for {
		attendancePack := dto.PackagerDTO{}

		receivedBody := <-provider.RequestBodyChan

		attendancePack.FirstID = receivedBody.Tag

		control := 0

		for {
			if control == 0 {
				control++
			} else {
				attendancePack.AttendanceID = append(attendancePack.AttendanceID, receivedBody.Tag)

			}

			receivedBody = <-provider.RequestBodyChan

			if receivedBody.Tag == attendancePack.FirstID {
				break
			}
		}

		provider.PackChan <- attendancePack
	}
}
