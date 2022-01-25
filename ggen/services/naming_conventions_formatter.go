package services

// NamingConventionsFormatter define las operaciones necesarias para formatear o transformar
// de un tipo de un tipo de convención a otro.
type NamingConventionsFormatter interface {
	ToPascalCase(name string) (string, error)
	ToCamelCase(name string) (string, error)
	ToKebabCase(name string) (string, error)
	ToSneakCase(name string) (string, error)
}
