package utils

import (
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/services"
	"github.com/stretchr/testify/assert"
)

/*
 funciones de ayuda
*/

// AssertSrcFiles recibe una lista de archivos de código fuente esperados y los compara
func AssertSrcFilesEqual(assert *assert.Assertions, expectedFiles []services.SrcFile, actualFiles []services.SrcFile) {
	//TODO assert varios de los archivos generados.
	for index, expectedFile := range expectedFiles {
		actualFile := actualFiles[index]
		assert.Equal(expectedFile, actualFile)
	}
}
