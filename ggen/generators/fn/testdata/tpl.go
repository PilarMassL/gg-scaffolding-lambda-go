package testdata

import "github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/models"

/*
 Archivos fuente rellenando usando las plantillas con datos de prueba.
*/

func ExpectedDummyReadMePopulated() models.SrcFile {
	return models.SrcFile{
		AbsolutePath: "./my_function_hello_world/readme.md",
		Content: `
# MyFunctionHelloWorld
Mi primera función usando el gg.

Esta función se disparará por un evento:
- evento-http
`,
	}
}
