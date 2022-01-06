package generators

// SrcTpl representa una plantilla de código fuente
type SrcTpl struct {
	RelativePath string //ruta absoluta del archivo
	Content      string //contenido del archivo
}

// SrcFile representa un archivo de código fuente.
type SrcFile struct {
	AbsolutePath string //ruta absoluta del archivo
	Content      string //contenido del archivo
}

// Generator define un objeto con la capacidad de generar archivos de código fuente.
type Generator interface {
	Generate(params interface{}) ([]SrcFile, error)
}

// GeneratorHelper define un objeto que permite realizar operaciones útiles de nombres para los generadores.
type GeneratorHelper interface {
	ValidateName(name string) string
	FillTpl(tpl SrcTpl, params interface{}, wd string) (SrcFile, error)
	FillTpls(tpl []SrcTpl, params interface{}, wd string) ([]SrcFile, error)
}

// Writer define un objeto con la capacidad de escribir en el disco archivos de código fuente generados.
type Writer interface {
	Save([]SrcFile) error
}
