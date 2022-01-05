package proyect

import (
	"errors"

	"github.com/PilarMassL/gg-scaffolding-lambda-go/poc-scaffolding-tool-go/ggen/generators"
)

// ProyectGenerator genera la estructura base de un proyecto al cuál se le pueden agregar funciones.
type ProyectGenerator struct {
	helper generators.GeneratorHelper
	wd     string
}

func NewProjectGenerator(helper generators.GeneratorHelper, wd string) *ProyectGenerator {
	return &ProyectGenerator{
		helper: helper,
		wd:     wd,
	}
}

func (g *ProyectGenerator) Generate(params generators.Params) ([]generators.SrcFile, error) {
	return nil, errors.New("No implementado")
}
