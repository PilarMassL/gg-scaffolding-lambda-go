package testdata

import "github.com/PilarMassL/gg-scaffolding-lambda-go/poc-scaffolding-tool-go/ggen/generators"

/*
 Archivos fuente rellando usando las plantillas con datos de prueba.
*/

func ExpectedReadMePopulated() generators.SrcFile {
	return generators.SrcFile{
		AbsolutePath: "./readme.md",
		Content: `
# MyProject
Mi primer proyecto usando el gg.
`,
	}
}
