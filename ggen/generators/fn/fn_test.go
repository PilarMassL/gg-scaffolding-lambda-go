package project

import (
	"testing"

	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/generators/project/testdata"
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/services"
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/utils"
	"github.com/stretchr/testify/assert"
)

// Debería generar los archivos indicados
func TestGenerateSuccess(t *testing.T) {
	//Arrange
	//El Writer agrega el efecto colateral, por eso se decide usar FakeWriter.
	writer := services.NewFakeWriter()
	generator := NewFnGenerator(writer)

	params := FnParams{
		FnName: "MyFunction",
	}

	//Act
	actualFiles, err := generator.Generate(params)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)
	if assert.NotNil(actualFiles) {
		expectedFiles := []services.SrcFile{
			testdata.ExpectedReadMePopulated(),
		}
		utils.AssertSrcFilesEqual(assert, expectedFiles, actualFiles)
	}
}
