package templates

import "github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/services"

func ReadMe() services.SrcTpl {
	return services.SrcTpl{
		RelativePath: "./readme.md",
		Content: `
# {{ .ProjectName }}
Mi primer proyecto usando el gg.
`,
	}
}
