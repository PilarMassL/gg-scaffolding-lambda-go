package generators

// SrcTpl representa una plantilla de código fuente
type SrcTpl struct {
	RelativePath string //ruta relativa del archivo
	Content      string //contenido del archivo
}

// SrcFile representa un archivo de código fuente.
type SrcFile struct {
	AbsolutePath string //ruta absoluta del archivo
	Content      string //contenido del archivo
}

// Generator define un objeto con la capacidad de generar archivos de código fuente.
type Generator interface {
	Generate(params interface{}) ([]SrcFile, error) //Generate = get templates + FillTplsAndSave (Efecto lateral)
}

// GeneratorHelper define un objeto que permite realizar operaciones útiles para construir generadores.
type GeneratorHelper interface {
	ValidateName(name string) string
	FillTpl(tpl SrcTpl, params interface{}) (SrcTpl, error)
	FillTpls(tpls []SrcTpl, params interface{}) ([]SrcTpl, error)                        //FillTpl sobre una lista
	FillTplsAndSave(tpls []SrcTpl, params interface{}, writer Writer) ([]SrcFile, error) //FillTpls + Save (Efecto lateral)
}

// Writer define un objeto con la capacidad de guardar archivos de código fuente generados.
type Writer interface {
	Save([]SrcTpl) ([]SrcFile, error)
}
