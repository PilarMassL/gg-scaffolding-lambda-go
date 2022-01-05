package generators

// Params representa los parametros que son reemplazados dentro de las plantillas.
type Params map[string]string

// SrcTpl representa una plantilla de código fuente
type SrcTpl struct {
	RelativePath string //ruta absoluta del archivo
	Content      []byte //contenido del archivo
}

// SrcFile representa un archivo de código fuente.
type SrcFile struct {
	AbsolutePath string //ruta absoluta del archivo
	Content      []byte //contenido del archivo
}

// Generator define un objeto con la capacidad de generar archivos de código fuente.
type Generator interface {
	Generate(params Params) ([]SrcFile, error)
}

// GeneratorHelper define un objeto que permite realizar operaciones útiles de nombres para los generadores.
type GeneratorHelper interface {
	ValidateName(name string) string
	FillTpl(tpl SrcTpl, params Params, wd string) (SrcFile, error)
}

// Writer define un objeto con la capacidad de escribir en el disco archivos de código fuente generados.
type Writer interface {
	Save([]SrcFile) error
}
