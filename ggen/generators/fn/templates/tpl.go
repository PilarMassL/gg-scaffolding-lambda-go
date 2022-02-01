package templates

import "github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/models"

func DummyReadMe() models.SrcTpl {
	return models.SrcTpl{
		RelativePath: "./{{ .FnDir }}/readme.md",
		Content: `
# {{ .FnName }}

Mi primera función usando el gg.

- Esta función se disparará por un evento: "{{ .FnEvent }}"
- Proyecto Base: "{{ .FnBaseProject }}"

`,
	}
}
