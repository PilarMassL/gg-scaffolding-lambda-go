package generators

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createDummySrcFile() []SrcFile {
	wd, _ := os.Getwd()
	return []SrcFile{
		{
			AbsolutePath: fmt.Sprintf("%s/test.txt", wd),
			Content:      `hola mundo`,
		},
	}
}

func TestSaveSuccess(t *testing.T) {
	//Arrange
	writer := NewGeneratorWriter()
	files := createDummySrcFile()

	//Act
	err := writer.Save(files)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)
}
