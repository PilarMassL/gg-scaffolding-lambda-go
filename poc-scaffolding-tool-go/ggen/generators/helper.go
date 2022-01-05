package generators

import (
	"errors"
)

// GeneratorHelper representa la implementación por defecto de un generator helper.
type GeneratorHelperDefault struct {
}

func NewGeneratorHelper() *GeneratorHelperDefault {
	return &GeneratorHelperDefault{}
}

//ValidateName permite validar si un nombre cumple con los requisitos necesarios para nombrar
//un elemento dentro de go.
func (g *GeneratorHelperDefault) ValidateName(name string) string {
	return ""
}

//FillTpl rellena una plantilla con los valores especificados en params
func (g *GeneratorHelperDefault) FillTpl(tpl SrcTpl, params Params, wd string) (SrcFile, error) {
	return SrcFile{}, errors.New("No implementado")
}
