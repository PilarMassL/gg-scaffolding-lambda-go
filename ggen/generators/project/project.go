package project

import (
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/generators/project/templates"
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/models"
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/services/generator"
)

// ProjectGenerator genera la estructura base de un proyecto al cuál se le pueden agregar funciones.
type ProjectGenerator struct {
	gen generator.GeneratorSvc
}

// ProjectParams contiene todos los parámetros necesarios para generar el proyecto base.
type ProjectPromptParams struct {
	ProjectName string `prompt:"¿Cuál es el nombre del proyecto?"`
}

func NewProjectGenerator(generator generator.GeneratorSvc) *ProjectGenerator {
	return &ProjectGenerator{
		// Acoplamos el Helper al generador, con la ventaja de simplificar la creación.
		gen: generator,
	}
}

func (p *ProjectGenerator) Generate(params ProjectPromptParams) ([]models.SrcFile, error) {
	//obtenemos las plantillas
	tpls := []models.SrcTpl{
		templates.ReadMe(),
	}
	//Rellenamos y guardamos
	return p.gen.FillTplsAndSave(tpls, params)
}
