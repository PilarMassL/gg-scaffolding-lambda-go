package naming_conventions

// Formatter define las operaciones necesarias para formatear o transformar
// de un tipo de un tipo de convención a otro.
type Formatter interface {
	ToPascalCase(name string) string
	ToCamelCase(name string) string
	ToKebabCase(name string) string
	ToSnakeCase(name string) string
}
