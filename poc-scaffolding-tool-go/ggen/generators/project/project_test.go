package proyect

import (
	"fmt"
	"testing"

	"github.com/PilarMassL/gg-scaffolding-lambda-go/poc-scaffolding-tool-go/ggen/generators"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*
  Test objects
*/

// MyMockedHelper es un mock de tipo GeneratorHelper
type MyMockedHelper struct {
	mock.Mock
}

// Se implementa la interfaz generators.GeneratorHelper
func (m *MyMockedHelper) FillTpl(tpl generators.SrcTpl, params generators.Params, wd string) (generators.SrcFile, error) {
	args := m.Called(tpl, params, wd)
	return args.Get(0).(generators.SrcFile), args.Error(1)
}

func (m *MyMockedHelper) ValidateName(name string) string {
	args := m.Called(name)
	return args.String(0)
}

// Debería generar los archivos indicados
func TestGeneratelSuccess(t *testing.T) {
	//Arrange
	wd := "/home/test"
	helper := new(MyMockedHelper)
	generator := NewProjectGenerator(helper, wd)

	params := make(generators.Params)
	params["projectName"] = "MyProject"

	// setup expectations
	helper.On("ValidateName", mock.Anything).Return("MyProject", nil)

	//Act
	files, err := generator.Generate(params)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)
	if assert.NotNil(files) {
		for index, file := range files {
			//TODO assert de los archivos generados.
			fmt.Println(index, "=> file.AbsolutePath: ", file.AbsolutePath)
		}
	}
}
