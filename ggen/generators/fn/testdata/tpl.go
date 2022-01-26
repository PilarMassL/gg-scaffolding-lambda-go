package testdata

import "github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/models"

/*
 Archivos fuente rellenando usando las plantillas con datos de prueba.
*/

func ExpectedDummyReadMePopulated() models.SrcFile {
	return models.SrcFile{
		AbsolutePath: "./my-function-hello-world/readme.md",
		Content: `
# MyFunctionHelloWorld
Mi primera función usando el gg.

A continuación mostramos el nombre de la función en distintas convenciones:

- Camel Case: myFunctionHelloWorld
- Pascal Case: MyFunctionHelloWorld
- Kebab Case: my-function-hello-world
- Snake Case: my_function_hello_world
`,
	}
}
