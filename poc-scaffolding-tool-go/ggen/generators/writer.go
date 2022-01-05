package generators

import (
	"errors"
)

// Validator define un objeto que permite realizar validaciones útiles de nombres para los generadores.
type GeneratorWriter struct {
}

func NewGeneratorWriter() *GeneratorWriter {
	return &GeneratorWriter{}
}

func (w *GeneratorWriter) Save(files []SrcFile) error {
	return errors.New("No implementado")
}
