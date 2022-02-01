package naming_conventions

import (
	"regexp"

	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/utils"
)

// ValidatorSvc representa la implementación por defecto de un
// Validator.
type ValidatorSvc struct {
	reCamelCaseValidator  *regexp.Regexp
	rePascalCaseValidator *regexp.Regexp
	reSnakeCaseValidator  *regexp.Regexp
	reKebabCaseValidator  *regexp.Regexp
}

func NewValidatorSvc() *ValidatorSvc {
	reCamelCaseValidator, errCCV := regexp.Compile(`[a-z]+((\d)|([A-Z0-9][a-z0-9]+))*([A-Z])?`)
	utils.CheckErrorFatal(errCCV)
	rePascalCaseValidator, errPCV := regexp.Compile(`([A-Z][a-z0-9]+)*([A-Z])?`)
	utils.CheckErrorFatal(errPCV)
	reSnakeCaseValidator, errSCV := regexp.Compile(`([a-z]+\_?[a-z0-9]+\_?)*`)
	utils.CheckErrorFatal(errSCV)
	reKebabCaseValidator, errKCV := regexp.Compile(`([a-z]+\-?[a-z0-9]+\-?)*`)
	utils.CheckErrorFatal(errKCV)

	return &ValidatorSvc{
		reCamelCaseValidator:  reCamelCaseValidator,
		rePascalCaseValidator: rePascalCaseValidator,
		reSnakeCaseValidator:  reSnakeCaseValidator,
		reKebabCaseValidator:  reKebabCaseValidator,
	}
}

func (v *ValidatorSvc) IsCamelCase(name string) bool {
	return checkNameWithRegex(name, v.reCamelCaseValidator)
}

func (v *ValidatorSvc) IsPascalCase(name string) bool {
	return checkNameWithRegex(name, v.rePascalCaseValidator)
}

func (v *ValidatorSvc) IsSnakeCase(name string) bool {
	return checkNameWithRegex(name, v.reSnakeCaseValidator)
}

func (v *ValidatorSvc) IsKebabCase(name string) bool {
	return checkNameWithRegex(name, v.reKebabCaseValidator)
}

//Funciones de utilidad

func checkNameWithRegex(name string, re *regexp.Regexp) bool {
	found := re.FindString(name)
	if found == "" || found != name {
		return false
	}
	return true
}
