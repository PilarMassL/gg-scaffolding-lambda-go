package generators

import (
	"bytes"
	"log"
	"text/template"
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
func (g *GeneratorHelperDefault) FillTpl(tpl SrcTpl, params interface{}) (SrcTpl, error) {

	var srcContent bytes.Buffer

	//Nota: es posible que ésta no sea la manera más eficiente de usar las plantillas.
	//Por el momento nos preocupamos del aspecto funcional y no tanto del rendimiento.

	//Se crea la plantilla.
	t, errCreatingTpl := template.New(tpl.RelativePath).Parse(tpl.Content)
	if errCreatingTpl != nil {
		return SrcTpl{}, errCreatingTpl
	}
	//Se mapean los parámetros
	errMappingTpl := t.Execute(&srcContent, params)
	if errMappingTpl != nil {
		return SrcTpl{}, errMappingTpl
	}

	srcTpl := SrcTpl{
		RelativePath: tpl.RelativePath,
		Content:      srcContent.String(),
	}
	return srcTpl, nil
}

// FillTpls rellena un arreglo de plantillas con los valores especificados en el argumento 'params'.
func (g *GeneratorHelperDefault) FillTpls(tpls []SrcTpl, params interface{}) ([]SrcTpl, error) {

	tplsFilled := make([]SrcTpl, 0, len(tpls))

	//Mapeamos las plantillas SrcFile útilizando FillTpl
	for _, tpl := range tpls {
		tplFilled, err := g.FillTpl(tpl, params)
		if err != nil {
			return nil, err
		}
		tplsFilled = append(tplsFilled, tplFilled)
		log.Printf("[Generator] plantilla '%s' rellenada correctamente.", tpl.RelativePath)
	}
	return tplsFilled, nil
}

// FillTplsAndSave rellena un arreglo de plantillas y lo guarda (EFECTO COLATERAL) usando un Writer
func (g *GeneratorHelperDefault) FillTplsAndSave(tpls []SrcTpl, params interface{}, writer Writer) ([]SrcFile, error) {
	// Rellenamos las plantillas.
	tplsFilled, errFillingTpls := g.FillTpls(tpls, params)
	if errFillingTpls != nil {
		return nil, errFillingTpls
	}

	// Guardamos los archivos generados.
	files, errSavingFiles := writer.Save(tplsFilled)
	if errSavingFiles != nil {
		return nil, errSavingFiles
	}
	return files, nil
}
