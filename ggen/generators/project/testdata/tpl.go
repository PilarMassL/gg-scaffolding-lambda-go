package testdata

import "github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/services"

/*
 Archivos fuente rellenando usando las plantillas con datos de prueba.
*/

func ExpectedReadMePopulated() services.SrcFile {
	return generators.SrcFile{
		AbsolutePath: "./readme.md",
		Content: `
# MyProject
Mi primer proyecto usando el gg.
`,
	}
}
