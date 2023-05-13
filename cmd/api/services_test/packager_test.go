package services_test

import (
	"testing"

	domain "github.com/FelipeMandelli/ProjetoInterativoV/cmd/api/internal/domain/rest"
	"github.com/FelipeMandelli/ProjetoInterativoV/cmd/api/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestPackager(t *testing.T) {
	// Arrange
	provider := services.GetProvider()
	body1 := domain.RequestBody{
		Tag: "123",
	}
	body2 := domain.RequestBody{
		Tag: "321",
	}
	body3 := domain.RequestBody{
		Tag: "456",
	}
	expectedStr := "123"
	expectedArr := []string{"321", "456"}

	// Act
	go services.Packager(provider)
	provider.RequestBodyChan <- body1
	provider.RequestBodyChan <- body2
	provider.RequestBodyChan <- body3
	provider.RequestBodyChan <- body1

	got := <-provider.PackChan

	// Assert
	assert.Equal(t, got.FirstID, expectedStr)
	assert.Equal(t, expectedArr, got.AttendanceID)
}

func TestMultiplePackager(t *testing.T) {
	// Arrange
	provider := services.GetProvider()
	body1 := domain.RequestBody{
		Tag: "123",
	}
	body2 := domain.RequestBody{
		Tag: "321",
	}
	body3 := domain.RequestBody{
		Tag: "456",
	}
	expectedStr1 := "123"
	expectedArr1 := []string{"321", "456"}
	expectedStr2 := "321"
	expectedArr2 := []string{"456", "456", "123"}

	// Act
	go services.Packager(provider)
	provider.RequestBodyChan <- body1
	provider.RequestBodyChan <- body2
	provider.RequestBodyChan <- body3
	provider.RequestBodyChan <- body1

	got1 := <-provider.PackChan

	provider.RequestBodyChan <- body2
	provider.RequestBodyChan <- body3
	provider.RequestBodyChan <- body3
	provider.RequestBodyChan <- body1
	provider.RequestBodyChan <- body2

	got2 := <-provider.PackChan

	// Assert
	assert.Equal(t, expectedStr1, got1.FirstID)
	assert.Equal(t, expectedArr1, got1.AttendanceID)
	assert.Equal(t, expectedStr2, got2.FirstID)
	assert.Equal(t, expectedArr2, got2.AttendanceID)
}
