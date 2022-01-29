package naming_conventions

import (
	"bytes"
	"regexp"
	"strings"

	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/utils"
)

// FormatterSvc representa la implementación por defecto de un
// Formatter.
type FormatterSvc struct {
	validator   Validator
	reKebabCase *regexp.Regexp
	reCamelCase *regexp.Regexp
}

func NewFormatterSvc(validator Validator) *FormatterSvc {
	reKebabCase, errKC := regexp.Compile(`(.*?)_([a-zA-Z])`)
	utils.CheckErrorFatal(errKC)
	reCamelCase, errCC := regexp.Compile(`(.+?)([A-Z])`)
	utils.CheckErrorFatal(errCC)

	return &FormatterSvc{
		validator:   validator,
		reKebabCase: reKebabCase,
		reCamelCase: reCamelCase,
	}
}

func (f *FormatterSvc) ToCamelCase(name string) string {
	if f.validator.IsPascalCase(name) {
		return lowerFirst(name)
	}
	if f.validator.IsSnakeCase(name) {
		return f.snake2Camel(name)
	}
	if f.validator.IsKebabCase(name) {
		// Kebab -> Snake -> Camel
		return f.snake2Camel(replace(name, "-", "_"))
	}
	return name
}

func (f *FormatterSvc) ToPascalCase(name string) string {
	if f.validator.IsCamelCase(name) {
		return upperFirst(name)
	}
	if f.validator.IsSnakeCase(name) {
		// Snake -> Camel -> Pascal
		return upperFirst(f.snake2Camel(name))
	}
	if f.validator.IsKebabCase(name) {
		// Kebab -> Snake -> Camel -> Pascal
		return upperFirst(f.snake2Camel(replace(name, "-", "_")))
	}
	return name
}

func (f *FormatterSvc) ToSnakeCase(name string) string {
	if f.validator.IsKebabCase(name) {
		return replace(name, "-", "_")
	}
	if f.validator.IsCamelCase(name) {
		return f.camel2Snake(name)
	}
	if f.validator.IsPascalCase(name) {
		// Pascal -> Camel -> Snake
		return f.camel2Snake(lowerFirst(name))
	}
	return replace(strings.ToLower(name), " ", "_")
}

func (f *FormatterSvc) ToKebabCase(name string) string {
	if f.validator.IsSnakeCase(name) {
		return replace(name, "_", "-")
	}
	if f.validator.IsCamelCase(name) {
		//Camel -> Snake -> Kebab
		return replace(f.camel2Snake(name), "_", "-")
	}
	if f.validator.IsPascalCase(name) {
		// Pascal -> Camel -> Snake -> Kebab
		return replace(f.camel2Snake(lowerFirst(name)), "_", "-")
	}
	return replace(strings.ToLower(name), " ", "-")
}

func (f *FormatterSvc) camel2Snake(name string) string {
	groups := f.reCamelCase.FindAllStringSubmatchIndex(name, -1)
	result := ""
	tailStartIndex := len(name)
	for _, g := range groups {
		g1 := name[g[2]:g[3]]
		g2 := name[g[4]:g[5]]
		tailStartIndex = g[5]
		result = result + g1 + "_" + strings.ToLower(g2)
	}
	result = result + name[tailStartIndex:]

	return result
}

func (f *FormatterSvc) snake2Camel(name string) string {
	groups := f.reKebabCase.FindAllStringSubmatchIndex(name, -1)
	result := ""
	tailStartIndex := len(name)
	for _, g := range groups {
		g1 := name[g[2]:g[3]]
		g2 := name[g[4]:g[5]]
		tailStartIndex = g[5]
		result = result + g1 + strings.ToUpper(g2)
	}
	result = result + name[tailStartIndex:]

	return result
}

//Funciones de utilidad

func replace(s, old, new string) string {
	return strings.Replace(s, old, new, -1)
}

func lowerFirst(s string) string {
	// Nota: s debe tener al menos 1 carácter.
	bts := []byte(s)

	fc := bytes.ToLower([]byte{bts[0]})
	rest := bts[1:]

	return string(bytes.Join([][]byte{fc, rest}, nil))
}

func upperFirst(s string) string {
	// Nota: s debe tener al menos 1 carácter.
	bts := []byte(s)

	fc := bytes.ToUpper([]byte{bts[0]})
	rest := bts[1:]

	return string(bytes.Join([][]byte{fc, rest}, nil))
}
