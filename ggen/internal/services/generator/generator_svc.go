package generator

import (
	"bytes"
	"log"
	"text/template"

	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/models"
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/services/writer"
)

// GeneratorSvc representa la implementación por defecto de un generator.
type GeneratorSvcDefault struct {
	writer writer.Writer
}

func NewGeneratorSvc(writer writer.Writer) *GeneratorSvcDefault {
	return &GeneratorSvcDefault{
		writer: writer,
	}
}

// FillTpl rellena una plantilla con los valores especificados en el argumento 'params'.
// Es importante que el campo 'tpl.Content' contenga los 'placeholders' en el formato:
// {{ .PlaceholderName }} y que la estructura 'params' tenga un campo llamado 'PlaceholderName'.
// Para más información visite: https://pkg.go.dev/text/template#Template
// Si no se cumple, es posible que encuentre errores difíciles de detectar, es RESPONSABILIDAD
// del desarrollador del generador garantizar ésta coherencia.
//
// También permite construir la ruta absoluta a partir de:
// - tpl.RelativePath: Ruta relativa del archivo fuente dentro del directorio de trabajo.
//   Normalmente debe tener el formato: './dir1/dir2/file.ext' (siempre arranca con punto)
func (g *GeneratorSvcDefault) FillTpl(tpl models.SrcTpl, params interface{}) (*models.SrcTpl, error) {

	//Nota: es posible que ésta no sea la manera más eficiente de usar las plantillas.
	//Por el momento nos preocupamos del aspecto funcional y no tanto del rendimiento.

	// el path también puede ser una plantilla
	relativePathFilled, err := fillTpl(tpl.RelativePath, tpl.RelativePath, params)
	if err != nil {
		return nil, err
	}

	contentFilled, err := fillTpl(relativePathFilled, tpl.Content, params)
	if err != nil {
		return nil, err
	}

	srcTpl := &models.SrcTpl{
		RelativePath: relativePathFilled,
		Content:      contentFilled,
	}
	return srcTpl, nil
}

// FillTpls rellena un arreglo de plantillas con los valores especificados en el argumento 'params'.
func (g *GeneratorSvcDefault) FillTpls(tpls []models.SrcTpl, params interface{}) ([]models.SrcTpl, error) {

	tplsFilled := make([]models.SrcTpl, 0, len(tpls))

	//Mapeamos las plantillas SrcFile útilizando FillTpl
	for _, tpl := range tpls {
		tplFilled, err := g.FillTpl(tpl, params)
		if err != nil {
			return nil, err
		}
		tplsFilled = append(tplsFilled, *tplFilled)
		log.Printf("[Generator] plantilla '%s' rellenada correctamente.", tplFilled.RelativePath)
	}
	return tplsFilled, nil
}

// FillTplsAndSave rellena un arreglo de plantillas y lo guarda (EFECTO COLATERAL) usando un Writer
func (g *GeneratorSvcDefault) FillTplsAndSave(tpls []models.SrcTpl, params interface{}) ([]models.SrcFile, error) {
	// Rellenamos las plantillas.
	tplsFilled, errFillingTpls := g.FillTpls(tpls, params)
	if errFillingTpls != nil {
		return nil, errFillingTpls
	}
	log.Print("[Generator] Plantillas rellenadas")

	// Guardamos los archivos generados.
	files, errSavingFiles := g.writer.Save(tplsFilled)
	if errSavingFiles != nil {
		return nil, errSavingFiles
	}
	log.Print("[Generator] Plantillas guardadas")
	return files, nil
}

func fillTpl(name, tpl string, params interface{}) (string, error) {
	var tplFilled bytes.Buffer
	t, err := template.New(name).Parse(tpl)
	if err != nil {
		return "", err
	}
	//Se mapean los parámetros
	errMappingTpl := t.Execute(&tplFilled, params)
	if errMappingTpl != nil {
		return "", errMappingTpl
	}
	return tplFilled.String(), nil

}
