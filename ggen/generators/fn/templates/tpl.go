package templates

import "github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/models"

func DummyReadMe() models.SrcTpl {
	return models.SrcTpl{
		RelativePath: "./my-function-hello-world/readme.md",
		Content: `
# {{ .FunctionNamePascalCase }}
Mi primera función usando el gg.

A continuación mostramos el nombre de la función en distintas convenciones:

- Camel Case: {{ .FunctionNameCamelCase }}
- Pascal Case: {{ .FunctionNamePascalCase }}
- Kebab Case: {{ .FunctionNameKebabCase }}
- Snake Case: {{ .FunctionNameSnakeCase }}
`,
	}
}
