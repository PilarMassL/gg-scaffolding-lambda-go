package project

import (
	"testing"

	"github.com/PilarMassL/gg-scaffolding-lambda-go/poc-scaffolding-tool-go/ggen/generators"
	"github.com/stretchr/testify/assert"
)

/*
  Test objects


// MyMockedHelper es un mock de tipo GeneratorHelper
type MyMockedHelper struct {
	mock.Mock
}

// Se implementa la interfaz generators.GeneratorHelper
func (m *MyMockedHelper) FillTpl(tpl generators.SrcTpl, params interface{}, wd string) (generators.SrcFile, error) {
	args := m.Called(tpl, params, wd)
	return args.Get(0).(generators.SrcFile), args.Error(1)
}

func (m *MyMockedHelper) ValidateName(name string) string {
	args := m.Called(name)
	return args.String(0)
}
*/

// Debería generar los archivos indicados
func TestGenerateSuccess(t *testing.T) {
	//Arrange
	wd := "/home/test"
	//helper := new(MyMockedHelper)
	helper := generators.NewGeneratorHelper()
	generator := NewProjectGenerator(helper, wd)

	params := ProjectParams{
		ProjectName: "MyProject",
	}

	// setup expectations
	// helper.On("FillTpl", mock.Anything).Return("MyProject", nil)

	//Act
	files, err := generator.Generate(params)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)
	if assert.NotNil(files) {
		expectedFiles := []generators.SrcFile{
			expectedReadMePopulated(),
		}
		assertSrcFilesEqual(assert, expectedFiles, files)
	}
}

/*
 funciones de ayuda
*/

// assertSrcFiles recibe una lista de archivos de código fuente esperados y los compara
func assertSrcFilesEqual(assert *assert.Assertions, expectedFiles []generators.SrcFile, files []generators.SrcFile) bool {
	//TODO assert varios de los archivos generados.
	return assert.Equal(true, true)
}

/*
 Archivos fuente rellando usando las plantillas con datos de prueba.
*/

func expectedReadMePopulated() generators.SrcFile {
	return generators.SrcFile{
		AbsolutePath: "/home/test/readme.md",
		Content: `
# MyProject
Mi primer proyecto usando el gg.
`,
	}
}
