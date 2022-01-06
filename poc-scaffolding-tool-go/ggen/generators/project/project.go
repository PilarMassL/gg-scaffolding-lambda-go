package project

import (
	"github.com/PilarMassL/gg-scaffolding-lambda-go/poc-scaffolding-tool-go/ggen/generators"
	"github.com/PilarMassL/gg-scaffolding-lambda-go/poc-scaffolding-tool-go/ggen/generators/project/templates"
)

// ProjectGenerator genera la estructura base de un proyecto al cuál se le pueden agregar funciones.
type ProjectGenerator struct {
	helper generators.GeneratorHelper
	wd     string
}

// ProjectParams contiene todos los parámetros necesarios para generar el proyecto base.
type ProjectParams struct {
	ProjectName string
}

func NewProjectGenerator(helper generators.GeneratorHelper, wd string) *ProjectGenerator {
	return &ProjectGenerator{
		helper: helper,
		wd:     wd,
	}
}

func (g *ProjectGenerator) Generate(params ProjectParams) ([]generators.SrcFile, error) {
	//obtenemos las plantillas
	tpls := []generators.SrcTpl{
		templates.ReadMe(),
	}
	return g.helper.FillTpls(tpls, params, g.wd)
}
