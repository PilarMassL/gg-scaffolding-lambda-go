package writer

import "github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/models"

// FakeWriter define un objeto que se comporta como un Writer pero no guarda nada. USAR SOLO PARA PRUEBAS.
type FakeWriter struct {
}

func NewFakeWriter() *FakeWriter {
	return &FakeWriter{}
}

func (w *FakeWriter) Save(tplsFilled []models.SrcTpl) ([]models.SrcFile, error) {
	filesSaved := make([]models.SrcFile, 0, len(tplsFilled))
	for _, tplFilled := range tplsFilled {
		filesSaved = append(filesSaved, models.SrcFile{
			AbsolutePath: tplFilled.RelativePath,
			Content:      tplFilled.Content,
		})
	}
	return filesSaved, nil
}
