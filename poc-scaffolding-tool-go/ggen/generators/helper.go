package generators

import (
	"bytes"
	"fmt"
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
// - wd: Directorio de trabajo, que es donde se van a guardar los archivos generados.
// - tpl.RelativePath: Ruta relativa del archivo fuente dentro del directorio de trabajo.
//   Normalmente debe tener el formato: './dir1/dir2/file.ext' (siempre arranca con punto)
func (g *GeneratorHelperDefault) FillTpl(tpl SrcTpl, params interface{}, wd string) (SrcFile, error) {

	var srcContent bytes.Buffer

	//Nota: es posible que ésta no sea la manera más eficiente de usar las plantillas.
	//Por el momento nos preocupamos del aspecto funcional y no tando del rendimiento.

	//Se crea la plantilla.
	t, errCreatingTpl := template.New(tpl.RelativePath).Parse(tpl.Content)
	if errCreatingTpl != nil {
		return SrcFile{}, errCreatingTpl
	}
	//Se mapean los parámetros
	errMappingTpl := t.Execute(&srcContent, params)
	if errMappingTpl != nil {
		return SrcFile{}, errMappingTpl
	}

	srcFile := SrcFile{
		//Lo primero que debemos hacer es qutar el punto inicial de la ruta relativa
		AbsolutePath: fmt.Sprintf("%s%s", wd, trimLeftChar(tpl.RelativePath)),
		Content:      srcContent.String(),
	}
	return srcFile, nil
}

// FillTpls rellena un arreglo de plantillas con los valores especificados en el argumento 'params'.
func (g *GeneratorHelperDefault) FillTpls(tpls []SrcTpl, params interface{}, wd string) ([]SrcFile, error) {

	files := make([]SrcFile, 0, len(tpls))

	//Mapeamos las plantillas SrcFile útilizando el helper
	for _, tpl := range tpls {
		file, err := g.FillTpl(tpl, params, wd)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	return files, nil
}

// trimLeftChar: remueve el primer carácter de una cádena de carácteres.
func trimLeftChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return s[:0]
}
