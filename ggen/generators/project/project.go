package project

import (
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/generators/project/templates"
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/services"
)

// ProjectGenerator genera la estructura base de un proyecto al cuál se le pueden agregar funciones.
type ProjectGenerator struct {
	helper services.GeneratorHelper
	writer services.Writer
}

// ProjectParams contiene todos los parámetros necesarios para generar el proyecto base.
type ProjectParams struct {
	ProjectName string `prompt:"¿Cuál es el nombre del proyecto?"`
}

func NewProjectGenerator(writer services.Writer) *ProjectGenerator {
	return &ProjectGenerator{
		// Acoplamos el Helper al generador, con la ventaja de simplificar la creación.
		helper: services.NewGeneratorHelper(),
		writer: writer,
	}
}

func (g *ProjectGenerator) Generate(params ProjectParams) ([]services.SrcFile, error) {
	//obtenemos las plantillas
	tpls := []services.SrcTpl{
		templates.ReadMe(),
	}
	//Rellenamos y guardamos
	return g.helper.FillTplsAndSave(tpls, params, g.writer)
}
