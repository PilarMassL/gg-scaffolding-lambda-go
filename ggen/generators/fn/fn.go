package project

import (
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/generators/project/templates"
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/services"
)

// FnGenerator genera la estructura base de una función a la cuál se le pueden agregar servicios, puertos y adaptadores.
type FnGenerator struct {
	helper services.GeneratorHelper
	writer services.Writer
}

// FnParams contiene todos los parámetros necesarios para generar el esqueleto de una función.
type FnParams struct {
	FnName string `prompt:"¿Cuál es el nombre de la función?"`
}

func NewFnGenerator(writer services.Writer) *FnGenerator {
	return &FnGenerator{
		// Acoplamos el Helper al generador, con la ventaja de simplificar la creación.
		helper: services.NewGeneratorHelper(),
		writer: writer,
	}
}

func (g *FnGenerator) Generate(params FnParams) ([]services.SrcFile, error) {
	//obtenemos las plantillas
	tpls := []services.SrcTpl{
		templates.ReadMe(),
	}
	//Rellenamos y guardamos
	return g.helper.FillTplsAndSave(tpls, params, g.writer)
}
