package testdata

import "github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/models"

/*
 Archivos fuente rellenando usando las plantillas con datos de prueba.
*/

func ExpectedReadMePopulated() models.SrcFile {
	return models.SrcFile{
		AbsolutePath: "./readme.md",
		Content: `
# MyProject

Mi primer proyecto usando el gg.

`,
	}
}
