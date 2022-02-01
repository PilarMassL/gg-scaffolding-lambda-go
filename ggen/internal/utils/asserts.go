package utils

import (
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/models"
	"github.com/stretchr/testify/assert"
)

/*
 funciones de ayuda para pruebas
*/

// AssertSrcFiles recibe una lista de archivos de código fuente esperados y los compara
func AssertSrcFilesEqual(assert *assert.Assertions, expectedFiles []models.SrcFile, actualFiles []models.SrcFile) {
	// assert varios de los archivos generados.
	for index, expectedFile := range expectedFiles {
		actualFile := actualFiles[index]
		assert.Equal(expectedFile, actualFile)
	}
}
