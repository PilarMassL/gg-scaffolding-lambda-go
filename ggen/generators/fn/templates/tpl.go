package templates

import "github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/services"

func DummyReadMe() services.SrcTpl {
	return services.SrcTpl{
		RelativePath: "./{{ .FunctionNameKebabCase}}/readme.md",
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
