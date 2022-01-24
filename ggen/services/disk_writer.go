package services

import (
	"io"
	"os"
	"path/filepath"
)

// DiskWriter define un objeto que permite escribir objetos SrcFile en un almacén durable.
type DiskWriter struct {
	wd string
}

func NewDiskWriter(workingDirectory string) *DiskWriter {
	return &DiskWriter{
		wd: workingDirectory,
	}
}

func (w *DiskWriter) Save(tplsFilled []SrcTpl) ([]SrcFile, error) {
	filesSaved := make([]SrcFile, 0, len(tplsFilled))
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

func updateFileAbsolutePath(tpl SrcTpl, wd string) SrcFile {
	//Lo primero que debemos hacer es qutar el punto inicial de la ruta relativa
	return SrcFile{
		AbsolutePath: filepath.Join(wd, trimLeftChar(tpl.RelativePath)),
		Content:      tpl.Content,
	}
}

// saveFile guarda un objeto SrcFile en el disco
func saveFile(file SrcFile) error {
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
