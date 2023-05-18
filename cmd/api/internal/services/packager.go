package services

import (
	dto "github.com/FelipeMandelli/ProjetoInterativoV/pkg/DTO"
)

// Packager packages attendance info in DTO
func Packager(provider *Provider) {
	for {
		attendancePack := dto.PackagerDTO{}

		receivedBody := <-provider.RequestBodyChan

		attendancePack.TeacherID = receivedBody.Tag

		control := 0

		for {
			if control == 0 {
				control++
			} else {
				attendancePack.AttendanceIDs = append(attendancePack.AttendanceIDs, receivedBody.Tag)

			}

			receivedBody = <-provider.RequestBodyChan

			if receivedBody.Tag == attendancePack.TeacherID {
				break
			}
		}

		provider.PackChan <- attendancePack
	}
}
