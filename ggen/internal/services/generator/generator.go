package generator

import (
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/models"
)

// Generator define un objeto con la capacidad de generar archivos de código fuente.
type Generator interface {
	Generate(params interface{}) ([]models.SrcFile, error) //Generate = get templates + FillTplsAndSave (Efecto lateral)
}

// GeneratorSvc define un objeto que permite realizar operaciones útiles para construir generadores.
type GeneratorSvc interface {
	FillTpl(tpl models.SrcTpl, params interface{}) (models.SrcTpl, error)
	FillTpls(tpls []models.SrcTpl, params interface{}) ([]models.SrcTpl, error)         //FillTpl sobre una lista
	FillTplsAndSave(tpls []models.SrcTpl, params interface{}) ([]models.SrcFile, error) //FillTpls + Save (Efecto lateral)
}
