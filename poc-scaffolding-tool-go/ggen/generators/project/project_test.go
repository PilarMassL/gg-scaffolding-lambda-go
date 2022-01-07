package project

import (
	"testing"

	"github.com/PilarMassL/gg-scaffolding-lambda-go/poc-scaffolding-tool-go/ggen/generators"
	"github.com/PilarMassL/gg-scaffolding-lambda-go/poc-scaffolding-tool-go/ggen/generators/project/testdata"
	"github.com/stretchr/testify/assert"
)

// Debería generar los archivos indicados
func TestGenerateSuccess(t *testing.T) {
	//Arrange
	//El Writer agrega el efecto colateral, por eso se decide usar FakeWriter.
	writer := generators.NewFakeWriter()
	generator := NewProjectGenerator(writer)

	params := ProjectParams{
		ProjectName: "MyProject",
	}

	//Act
	actualFiles, err := generator.Generate(params)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)
	if assert.NotNil(actualFiles) {
		expectedFiles := []generators.SrcFile{
			testdata.ExpectedReadMePopulated(),
		}
		assertSrcFilesEqual(assert, expectedFiles, actualFiles)
	}
}

/*
 funciones de ayuda
*/

// assertSrcFiles recibe una lista de archivos de código fuente esperados y los compara
func assertSrcFilesEqual(assert *assert.Assertions, expectedFiles []generators.SrcFile, actualFiles []generators.SrcFile) {
	//TODO assert varios de los archivos generados.
	for index, expectedFile := range expectedFiles {
		actualFile := actualFiles[index]
		assert.Equal(expectedFile, actualFile)
	}
}
