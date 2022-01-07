package generators

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*
  Test objects
*/

type DummyParams struct {
	Titulo string
}

// MyMockedWriter es un mock de tipo Writer
type MyMockedWriter struct {
	mock.Mock
}

// Se implementa la interfaz generators.Writer
func (m *MyMockedWriter) Save(tplsFilled []SrcTpl) ([]SrcFile, error) {
	args := m.Called(tplsFilled)
	return args.Get(0).([]SrcFile), args.Error(1)
}

/*
 Funciones de ayuda
*/

func createDummyParams(titulo string) *DummyParams {
	return &DummyParams{
		Titulo: titulo,
	}
}

func createDummySrcTpl() SrcTpl {
	return SrcTpl{
		RelativePath: "./doc/readme.md",
		Content:      `# {{ .Titulo }} Hola mundo`,
	}
}

// Debería llenar un arreglo de plantillas con los parámetros especificados, la longitud del arreglo de salida
// debe ser igual al del arreglo de entrada.
func TestFillTplsSuccess(t *testing.T) {
	//Arrange
	helper := NewGeneratorHelper()
	tpls := []SrcTpl{
		createDummySrcTpl(),
		{
			RelativePath: "./doc/LICENSE",
			Content:      `LICENSE dummy`,
		},
	}

	params := createDummyParams("Titulo:")

	//Act
	files, err := helper.FillTpls(tpls, params)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)
	if assert.NotNil(files) {
		assert.Equal(2, len(files))
	}
}

// Debería llenar una plantilla con los parámetros especificados
func TestFillTplSuccess(t *testing.T) {
	//Arrange
	helper := NewGeneratorHelper()
	tpl := createDummySrcTpl()
	params := createDummyParams("Titulo:")

	//Act
	file, err := helper.FillTpl(tpl, params)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)
	if assert.NotNil(file) {
		assert.Equal(`# Titulo: Hola mundo`, file.Content)
	}
}

//Debería fallar porque la plantilla no tiene el formato adecuado.
func TestFillTplFailureBadTemplate(t *testing.T) {
	//Arrange
	helper := NewGeneratorHelper()
	params := createDummyParams("")

	tpl := SrcTpl{
		RelativePath: "./doc/LICENSE",
		Content:      `LICENSE dummy {{ badPlaceholder }}`, //El formato correcto debería ser: .BadPlaceHolder
	}

	//Act
	_, err := helper.FillTpl(tpl, params)

	//Assert
	assert := assert.New(t)
	assert.NotNil(err)

}

//Debería fallar porque al menos una plantilla no tiene el formato adecuado.
func TestFillTplsFailureBadTemplate(t *testing.T) {
	//Arrange
	helper := NewGeneratorHelper()
	params := createDummyParams("")

	tpls := []SrcTpl{
		createDummySrcTpl(),
		{
			RelativePath: "./doc/LICENSE",
			Content:      `LICENSE dummy {{ badPlaceHolder }}`, //El formato correcto debería ser: .BadPlaceHolder
		},
	}

	//Act
	files, err := helper.FillTpls(tpls, params)

	//Assert
	assert := assert.New(t)
	assert.NotNil(err)
	assert.Nil(files)

}

//Debería ejecutar FillTpls y llamar el método Save del writer.
func TestFillTplsAndSave(t *testing.T) {
	//Arrange
	helper := NewGeneratorHelper()
	writer := new(MyMockedWriter)
	params := createDummyParams("")

	tpls := []SrcTpl{
		createDummySrcTpl(),
		{
			RelativePath: "./doc/LICENSE",
			Content:      `LICENSE dummy`,
		},
	}

	//set expectations
	writer.On("Save", mock.Anything).Return([]SrcFile{
		{
			AbsolutePath: "/home/test/doc/LICENSE",
			Content:      `LICENSE dummy`,
		},
	}, nil)

	//Act
	files, err := helper.FillTplsAndSave(tpls, params, writer)

	//Assert
	assert := assert.New(t)
	writer.AssertCalled(t, "Save", mock.Anything)
	assert.NotNil(files)
	assert.Nil(err)
}
