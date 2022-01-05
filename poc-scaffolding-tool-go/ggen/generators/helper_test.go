package generators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createDummySrcTpl() SrcTpl {
	return SrcTpl{
		RelativePath: "./doc/readme.md",
		Content:      []byte(`# {{ title }} Hola mundo`),
	}
}

// Debería llenar una plantilla con los parámetros especificados
func TestFillTplSuccess(t *testing.T) {
	//Arrange
	helper := NewGeneratorHelper()
	tpl := createDummySrcTpl()
	params := make(Params)
	params["title"] = "Titulo"
	wd := "/home/test"

	//Act
	file, err := helper.FillTpl(tpl, params, wd)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)
	if assert.NotNil(file) {
		assert.Equal([]byte(`# {{ Titulo }} Hola mundo`), file.Content)
		assert.Equal("/home/test/doc/readme.md", file.AbsolutePath)
	}
}
