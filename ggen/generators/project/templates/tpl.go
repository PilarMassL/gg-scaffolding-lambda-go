package templates

import "github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/generators"

func ReadMe() generators.SrcTpl {
	return generators.SrcTpl{
		RelativePath: "./readme.md",
		Content: `
# {{ .ProjectName }}
Mi primer proyecto usando el gg.
`,
	}
}
