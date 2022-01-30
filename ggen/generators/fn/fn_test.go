package fn

import (
	"testing"

	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/generators/fn/testdata"
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/models"
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/services/generator"
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/services/writer"
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/utils"
	"github.com/stretchr/testify/assert"
)

// Debería generar los archivos indicados
func TestGenerateSuccess(t *testing.T) {
	//Arrange
	//El Writer agrega el efecto colateral, por eso se decide usar FakeWriter.
	writer := writer.NewFakeWriter()
	base := generator.NewGeneratorSvc(writer)
	generator := NewFnGenerator(base)

	params := FnPromptParams{
		FnName: "MyFunctionHelloWorld",
	}

	//Act
	actualFiles, err := generator.Generate(params)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)
	if assert.NotNil(actualFiles) {
		expectedFiles := []models.SrcFile{
			testdata.ExpectedDummyReadMePopulated(),
		}
		utils.AssertSrcFilesEqual(assert, expectedFiles, actualFiles)
	}
}
