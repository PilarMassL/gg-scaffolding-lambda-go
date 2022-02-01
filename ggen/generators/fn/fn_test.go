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
		FnName:        "MyFunctionHelloWorld",
		FnBaseProject: "github.com/example",
		FnEvent:       Http,
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

//Debería devolver el FnEvents (Enumeración) a partir de un String
func TestParseFnEventsSuccess(t *testing.T) {
	//Arrange
	genericStr := "generic"
	httpStr := "http"
	s3Str := "s3"

	//Act
	genericResult, _ := ParseEvent(genericStr)
	httpResult, _ := ParseEvent(httpStr)
	s3Result, _ := ParseEvent(s3Str)

	//Assert
	assert := assert.New(t)
	assert.Equal(Generic, genericResult)
	assert.Equal(Http, httpResult)
	assert.Equal(S3, s3Result)
}

//Debería fallar al interpretar un string que no coincide con un FnEvents (Enumeración)
func TestParseFnEventsFailed(t *testing.T) {
	//Arrange
	inputStr := "k.e.s.o"

	//Act
	result, err := ParseEvent(inputStr)

	//Assert
	assert := assert.New(t)
	assert.NotNil(err)
	assert.Equal(Invalid, result)
}

func TestFnEventsToStringSuccess(t *testing.T) {
	//Arrange

	//Act
	genericResult := Generic.String()
	httpResult := Http.String()
	s3Result := S3.String()

	//Assert
	assert := assert.New(t)
	assert.Equal("generic", genericResult)
	assert.Equal("http", httpResult)
	assert.Equal("S3", s3Result)
}
