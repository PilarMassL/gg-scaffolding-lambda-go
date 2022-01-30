package writer

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/models"
)

// DiskWriter define un objeto que permite escribir objetos models.SrcFile en un almacén durable.
type DiskWriter struct {
	wd string
}

func NewDiskWriter(workingDirectory string) *DiskWriter {
	return &DiskWriter{
		wd: workingDirectory,
	}
}

func (w *DiskWriter) Save(tplsFilled []models.SrcTpl) ([]models.SrcFile, error) {
	filesSaved := make([]models.SrcFile, 0, len(tplsFilled))
	for _, tplFilled := range tplsFilled {
		fileWithAbsolutePathUpdated := updateFileAbsolutePath(tplFilled, w.wd)
		err := saveFile(fileWithAbsolutePathUpdated)
		if err != nil {
			return nil, err
		}
		filesSaved = append(filesSaved, fileWithAbsolutePathUpdated)
	}
	return filesSaved, nil
}

func updateFileAbsolutePath(tpl models.SrcTpl, wd string) models.SrcFile {
	//Lo primero que debemos hacer es quitar el punto inicial de la ruta relativa
	return models.SrcFile{
		AbsolutePath: filepath.Join(wd, tpl.RelativePath),
		Content:      tpl.Content,
	}
}

// saveFile guarda un objeto models.SrcFile en el disco
func saveFile(file models.SrcFile) error {
	log.Printf("[DiskWriter] Intentamos guardar el archivo en: %s", file.AbsolutePath)
	f, err := create(file.AbsolutePath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.WriteString(f, file.Content)
	return err
}

//Forma estándar: https://stackoverflow.com/questions/59961510/golang-os-create-path-with-nested-directories#59961623
func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}
