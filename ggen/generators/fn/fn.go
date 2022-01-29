package project

import (
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/generators/fn/templates"
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/models"
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/services/generator"
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/services/naming_conventions"
)

// FnGenerator genera la estructura base de una función a la cuál se le pueden agregar servicios, puertos y adaptadores.
type FnGenerator struct {
	gen       generator.GeneratorSvc
	validator naming_conventions.Validator
	formatter naming_conventions.Formatter
}

// FnEvents es una enumeración que indica todos los posibles eventos que pueden disparar una función
type FnEvents int

const (
	HttpEvent FnEvents = iota
	S3Event
)

func (fe FnEvents) String() string {
	return []string{"Evento Http", "Evento desde S3"}[fe]
}

// FnParams contiene todos los parámetros necesarios para generar el esqueleto de una función.
type FnPromptParams struct {
	FnName  string   `prompt:"¿Cuál es el nombre de la función?" type:"input"`
	FnEvent FnEvents `prompt:"¿Cuál es el tipo de evento que dispara la función?" type:"list"`
}

type FnParams struct {
	FnName  string
	FnEvent string
	FnDir   string
}

func NewFnGenerator(generatorBase generator.GeneratorSvc) *FnGenerator {
	validator := naming_conventions.NewValidatorSvc()
	return &FnGenerator{
		// Acoplamos el Helper al generador, con la ventaja de simplificar la creación.
		gen:       generatorBase,
		validator: validator,
		formatter: naming_conventions.NewFormatterSvc(validator),
	}
}

func (fn *FnGenerator) Generate(params FnPromptParams) ([]models.SrcFile, error) {
	//Es importante que al agregar una función se agreguen solo las dependencias necesarias.
	//Solo se pueden agregar funciones a proyectos ya iniciados.

	//obtenemos las plantillas
	tpls := []models.SrcTpl{
		templates.DummyReadMe(),
	}

	//Obtenemos los parámetros para la plantilla de los parametros del usuario
	tplParams := fn.completeTplParams(params)

	//Rellenamos y guardamos
	return fn.gen.FillTplsAndSave(tpls, tplParams)
}

func (fn *FnGenerator) completeTplParams(params FnPromptParams) *FnParams {
	return &FnParams{
		FnName:  fn.formatter.ToPascalCase(params.FnName),
		FnEvent: fn.formatter.ToKebabCase(params.FnEvent.String()),
		FnDir:   fn.formatter.ToSnakeCase(params.FnName),
	}
}
