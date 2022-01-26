package models

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
