package services

// NamingConventionsValidator define las operaciones para validar si una determinada cadena
// cumple con alguno de los tipo de convenciones de nombramiento.
type NamingConventionsValidator interface {
	IsCamelCase(name string) bool
	IsPascalCase(nae string) bool
	IsKebabCase(name string) bool
	IsSnakeCase(name string) bool
}
