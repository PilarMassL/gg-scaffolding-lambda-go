package generators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 Funciones de ayuda
*/

type DummyParams struct {
	Titulo string
}

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
			RelativePath: "./doc/LICENCE",
			Content:      `LICENCE dummy`,
		},
	}

	params := createDummyParams("Titulo:")
	wd := "/home/test"

	//Act
	files, err := helper.FillTpls(tpls, params, wd)

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
	wd := "/home/test"

	//Act
	file, err := helper.FillTpl(tpl, params, wd)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)
	if assert.NotNil(file) {
		assert.Equal(`# Titulo: Hola mundo`, file.Content)
	}
}

func TestBuildAbsolutePath(t *testing.T) {
	//Arrange
	helper := NewGeneratorHelper()
	tpl := createDummySrcTpl()
	params := createDummyParams("")
	wd := "/home/test"

	//Act
	file, err := helper.FillTpl(tpl, params, wd)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)
	if assert.NotNil(file) {
		assert.Equal("/home/test/doc/readme.md", file.AbsolutePath)
	}
}
