package project

import (
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/generators"
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/generators/project/templates"
)

// ProjectGenerator genera la estructura base de un proyecto al cuál se le pueden agregar funciones.
type ProjectGenerator struct {
	helper generators.GeneratorHelper
	writer generators.Writer
}

// ProjectParams contiene todos los parámetros necesarios para generar el proyecto base.
type ProjectParams struct {
	ProjectName string `prompt:"¿Cuál es el nombre del proyecto?"`
}

func NewProjectGenerator(writer generators.Writer) *ProjectGenerator {
	return &ProjectGenerator{
		// Acoplamos el Helper al generador, con la ventaja de simplificar la creación.
		helper: generators.NewGeneratorHelper(),
		writer: writer,
	}
}

func (g *ProjectGenerator) Generate(params ProjectParams) ([]generators.SrcFile, error) {
	//obtenemos las plantillas
	tpls := []generators.SrcTpl{
		templates.ReadMe(),
	}
	//Rellenamos y guardamos
	return g.helper.FillTplsAndSave(tpls, params, g.writer)
}
