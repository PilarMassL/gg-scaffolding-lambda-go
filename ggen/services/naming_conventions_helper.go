package services

import (
	"fmt"
	"os"
	"regexp"
)

// NamingConventionsHelper representa la implementación por defecto de un
// NamingConventionsValidator, NamingConventionsFormatter.
type NamingConventionsHelper struct {
	reCamelCase  *regexp.Regexp
	rePascalCase *regexp.Regexp
	reSnakeCase  *regexp.Regexp
	reKebabCase  *regexp.Regexp
}

func NewNamingConventionsHelper() *NamingConventionsHelper {
	reCamelCase, errCC := regexp.Compile(`[a-z]+((\d)|([A-Z0-9][a-z0-9]+))*([A-Z])?`)
	checkError(errCC)
	rePascalCase, errPC := regexp.Compile(`(([A-Z0-9][a-z0-9]+))*([A-Z])?`)
	checkError(errPC)
	reSnakeCase, errSC := regexp.Compile(`(([A-Z0-9][a-z0-9]+))*([A-Z])?`) //TODO
	checkError(errSC)
	reKebabCase, errKC := regexp.Compile(`(([A-Z0-9][a-z0-9]+))*([A-Z])?`) //TODO
	checkError(errKC)
	return &NamingConventionsHelper{
		reCamelCase:  reCamelCase,
		rePascalCase: rePascalCase,
		reSnakeCase:  reSnakeCase,
		reKebabCase:  reKebabCase,
	}
}

func (nc *NamingConventionsHelper) IsCamelCase(name string) bool {
	return checkNameWithRegex(name, nc.reCamelCase)
}

func (nc *NamingConventionsHelper) IsPascalCase(name string) bool {
	return checkNameWithRegex(name, nc.rePascalCase)
}

func (nc *NamingConventionsHelper) IsSnakeCase(name string) bool {
	return checkNameWithRegex(name, nc.reSnakeCase)
}

func (nc *NamingConventionsHelper) IsKebabCase(name string) bool {
	return checkNameWithRegex(name, nc.reKebabCase)
}

func (nc *NamingConventionsHelper) ToCamelCase(name string) (string, error) {
	return name, nil
}

func (nc *NamingConventionsHelper) ToPascalCase(name string) (string, error) {
	return name, nil
}

func (nc *NamingConventionsHelper) ToKebabCase(name string) (string, error) {
	return name, nil
}

func (nc *NamingConventionsHelper) ToSnakeCase(name string) (string, error) {
	return name, nil
}

func checkNameWithRegex(name string, re *regexp.Regexp) bool {
	found := re.FindString(name)
	if found == "" || found != name {
		return false
	}
	return true
}
func checkError(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
