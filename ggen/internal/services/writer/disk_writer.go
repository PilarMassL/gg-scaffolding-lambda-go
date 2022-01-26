package writer

import (
	"io"
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
	//Lo primero que debemos hacer es qutar el punto inicial de la ruta relativa
	return models.SrcFile{
		AbsolutePath: filepath.Join(wd, trimLeftChar(tpl.RelativePath)),
		Content:      tpl.Content,
	}
}

// saveFile guarda un objeto models.SrcFile en el disco
func saveFile(file models.SrcFile) error {
	f, err := os.Create(file.AbsolutePath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.WriteString(f, file.Content)
	if err != nil {
		return err
	}
	return nil
}

// trimLeftChar: remueve el primer carácter de una cadena de caracteres.
func trimLeftChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return s[:0]
}
