package templates

import "github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/models"

func ReadMe() models.SrcTpl {
	return models.SrcTpl{
		RelativePath: "./readme.md",
		Content: `
# {{ .ProjectName }}
Mi primer proyecto usando el gg.
`,
	}
}
