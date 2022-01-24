package services

// FakeWriter define un objeto que se comporta como un Writer pero no guarda nada. USAR SOLO PARA PRUEBAS.
type FakeWriter struct {
}

func NewFakeWriter() *FakeWriter {
	return &FakeWriter{}
}

func (w *FakeWriter) Save(tplsFilled []SrcTpl) ([]SrcFile, error) {
	filesSaved := make([]SrcFile, 0, len(tplsFilled))
	for _, tplFilled := range tplsFilled {
		filesSaved = append(filesSaved, SrcFile{
			AbsolutePath: tplFilled.RelativePath,
			Content:      tplFilled.Content,
		})
	}
	return filesSaved, nil
}
