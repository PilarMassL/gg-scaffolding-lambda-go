package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createFakeDummyTplFilled() []SrcTpl {
	return []SrcTpl{
		{
			RelativePath: "./test.txt",
			Content:      `hola mundo`,
		},
	}
}

// Debería simplemente devolver la lista de archivos indicada.
func TestFakeSaveSuccess(t *testing.T) {
	//Arrange
	writer := NewFakeWriter()
	tpls := createDummyTplFilled()

	//Act
	filesSaved, err := writer.Save(tpls)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)

	if assert.NotNil(filesSaved) {

		assert.Equal(len(tpls), len(filesSaved))
		// Validamos que el contenido sea el mismo.
		for index, file := range filesSaved {
			assert.Equal(tpls[index].Content, file.Content)
		}
	}
}
