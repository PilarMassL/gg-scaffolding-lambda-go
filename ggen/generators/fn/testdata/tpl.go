package testdata

import "github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/services"

/*
 Archivos fuente rellenando usando las plantillas con datos de prueba.
*/

func ExpectedDummyReadMePopulated() services.SrcFile {
	return services.SrcFile{
		AbsolutePath: "./my-function/readme.md",
		Content: `
# MyFunction
Mi primera función usando el gg.

A continuación mostramos el nombre de la función en distintas convenciones:

- Camel Case: myFunction
- Pascal Case: MyFunction
- Kebab Case: my-function
- Snake Case: my_function
`,
	}
}
