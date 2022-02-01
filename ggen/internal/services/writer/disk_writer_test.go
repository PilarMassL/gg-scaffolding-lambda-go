package writer

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/models"
)

func createDummyTplFilled() []models.SrcTpl {
	return []models.SrcTpl{
		{
			RelativePath: "./test.txt",
			Content:      `hola mundo`,
		},
	}
}

// Debería guardar en disco la lista de archivos indicada.
func TestDiskSaveSuccess(t *testing.T) {
	//Arrange
	//Se crea una carpeta temporal
	wd, err := os.MkdirTemp("", "template_test")
	if err != nil {
		log.Fatal(err)
	}

	writer := NewDiskWriter(wd)
	tpls := createDummyTplFilled()

	//Act
	filesSaved, err := writer.Save(tpls)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)

	if assert.NotNil(filesSaved) {

		assert.Equal(len(tpls), len(filesSaved))
		//Validamos que existan en disco los archivos
		for _, file := range filesSaved {
			assertFileExist(assert, file)
		}
	}
}

// Debería actualizar el campo 'AbsolutePath' de los objetos models.SrcFile con la ruta en la que fue almacenado.
func TestUpdateAbsolutePath(t *testing.T) {
	//Arrange
	//Se crea una carpeta temporal
	wd, err := os.MkdirTemp("", "template_test")
	if err != nil {
		log.Fatal(err)
	}

	writer := NewDiskWriter(wd)

	dummyTpls := []models.SrcTpl{
		{
			RelativePath: "./test.txt",
			Content:      `hola mundo`,
		},
	}
	expectedFiles := []models.SrcFile{
		{
			AbsolutePath: filepath.Join(wd, "test.txt"),
			Content:      `hola mundo`,
		},
	}

	//Act
	actualFiles, err := writer.Save(dummyTpls)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)
	if assert.NotNil(actualFiles) {
		for index, expectedFile := range expectedFiles {
			actualFile := actualFiles[index]
			assert.Equal(expectedFile, actualFile)
		}

	}
}

// Debería fallar al intentar guardar en disco.
func TestDiskSaveFailed(t *testing.T) {
	//Arrange
	//Se crea una carpeta temporal
	wd, err := os.MkdirTemp("", "template_test")
	if err != nil {
		log.Fatal(err)
	}

	writer := NewDiskWriter(wd)
	tpls := []models.SrcTpl{
		{
			RelativePath: "",
			Content:      `hola mundo`,
		},
	}

	//Act
	filesSaved, err := writer.Save(tpls)

	//Assert
	assert := assert.New(t)
	assert.NotNil(err)
	assert.Nil(filesSaved)
}

/*
 funciones de ayuda
*/

func assertFileExist(assert *assert.Assertions, srcFile models.SrcFile) {
	_, err := os.Open(srcFile.AbsolutePath)
	assert.False(errors.Is(err, os.ErrNotExist))
	//Vale la pena comparar el contenido guardado con el contenido del srcFile?
}
