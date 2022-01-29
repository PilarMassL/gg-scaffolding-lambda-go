package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/models"
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
func (m *MyMockedWriter) Save(tplsFilled []models.SrcTpl) ([]models.SrcFile, error) {
	args := m.Called(tplsFilled)
	return args.Get(0).([]models.SrcFile), args.Error(1)
}

/*
 Funciones de ayuda
*/

func createDummyParams(titulo string) *DummyParams {
	return &DummyParams{
		Titulo: titulo,
	}
}

func createDummySrcTpl() models.SrcTpl {
	return models.SrcTpl{
		RelativePath: "./doc/readme.md",
		Content:      `# {{ .Titulo }} Hola mundo`,
	}
}

// Debería llenar un arreglo de plantillas con los parámetros especificados, la longitud del arreglo de salida
// debe ser igual al del arreglo de entrada.
func TestFillTplsSuccess(t *testing.T) {
	//Arrange
	writer := new(MyMockedWriter)
	gen := NewGeneratorSvc(writer)

	tpls := []models.SrcTpl{
		createDummySrcTpl(),
		{
			RelativePath: "./doc/LICENSE",
			Content:      `LICENSE dummy`,
		},
	}

	params := createDummyParams("Titulo:")

	//Act
	files, err := gen.FillTpls(tpls, params)

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
	writer := new(MyMockedWriter)
	gen := NewGeneratorSvc(writer)

	tpl := createDummySrcTpl()
	params := createDummyParams("Titulo:")

	//Act
	file, err := gen.FillTpl(tpl, params)

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
	writer := new(MyMockedWriter)
	gen := NewGeneratorSvc(writer)

	params := createDummyParams("")

	tpl := models.SrcTpl{
		RelativePath: "./doc/LICENSE",
		Content:      `LICENSE dummy {{ badPlaceholder }}`, //El formato correcto debería ser: .BadPlaceHolder
	}

	//Act
	tplFilled, err := gen.FillTpl(tpl, params)

	//Assert
	assert := assert.New(t)
	assert.NotNil(err)
	assert.Nil(tplFilled)

}

//Debería fallar porque al menos una plantilla no tiene el formato adecuado.
func TestFillTplsFailureBadTemplate(t *testing.T) {
	//Arrange
	writer := new(MyMockedWriter)
	gen := NewGeneratorSvc(writer)

	params := createDummyParams("")

	tpls := []models.SrcTpl{
		createDummySrcTpl(),
		{
			RelativePath: "./doc/LICENSE",
			Content:      `LICENSE dummy {{ badPlaceHolder }}`, //El formato correcto debería ser: .BadPlaceHolder
		},
	}

	//Act
	files, err := gen.FillTpls(tpls, params)

	//Assert
	assert := assert.New(t)
	assert.NotNil(err)
	assert.Nil(files)

}

//Debería ejecutar FillTpls y llamar el método Save del writer.
func TestFillTplsAndSave(t *testing.T) {
	//Arrange

	writer := new(MyMockedWriter)
	gen := NewGeneratorSvc(writer)
	params := createDummyParams("")

	tpls := []models.SrcTpl{
		createDummySrcTpl(),
		{
			RelativePath: "./doc/LICENSE",
			Content:      `LICENSE dummy`,
		},
	}

	//set expectations
	writer.On("Save", mock.Anything).Return([]models.SrcFile{
		{
			AbsolutePath: "/home/test/doc/LICENSE",
			Content:      `LICENSE dummy`,
		},
	}, nil)

	//Act
	files, err := gen.FillTplsAndSave(tpls, params)

	//Assert
	assert := assert.New(t)
	writer.AssertCalled(t, "Save", mock.Anything)
	assert.NotNil(files)
	assert.Nil(err)
}
