package naming_conventions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var validator = NewValidatorSvc()

//ToCamelCase casos

func TestToCamelCaseFromCamelCase(t *testing.T) {
	//Arrange
	formatter := NewFormatterSvc(validator)
	pascalCaseName := "functionName"

	//Act
	actual := formatter.ToCamelCase(pascalCaseName)

	//Assert
	assert := assert.New(t)
	assert.Equal("functionName", actual)
}

func TestToCamelCaseFromPascalCase(t *testing.T) {
	//Arrange
	formatter := NewFormatterSvc(validator)
	pascalCaseName := "FunctionName"

	//Act
	actual := formatter.ToCamelCase(pascalCaseName)

	//Assert
	assert := assert.New(t)
	assert.Equal("functionName", actual)
}

func TestToCamelCaseFromSnakeCase(t *testing.T) {
	//Arrange
	formatter := NewFormatterSvc(validator)
	snakeCaseName := "my_function_name"

	//Act
	actual := formatter.ToCamelCase(snakeCaseName)

	//Assert
	assert := assert.New(t)

	assert.Equal("myFunctionName", actual)
}

func TestToCamelCaseFromKebabCase(t *testing.T) {
	//Arrange
	formatter := NewFormatterSvc(validator)
	kebabCaseName := "function-name"

	//Act
	actual := formatter.ToCamelCase(kebabCaseName)

	//Assert
	assert := assert.New(t)

	assert.Equal("functionName", actual)
}

//ToPascalCase casos
func TestToPascalCaseFromPascalCase(t *testing.T) {
	//Arrange
	formatter := NewFormatterSvc(validator)
	camelCaseName := "FunctionName"

	//Act
	actual := formatter.ToPascalCase(camelCaseName)

	//Assert
	assert := assert.New(t)

	assert.Equal("FunctionName", actual)
}

func TestToPascalCaseFromCamelCase(t *testing.T) {
	//Arrange
	formatter := NewFormatterSvc(validator)
	camelCaseName := "functionName"

	//Act
	actual := formatter.ToPascalCase(camelCaseName)

	//Assert
	assert := assert.New(t)

	assert.Equal("FunctionName", actual)
}

func TestToPascalCaseFromSnakeCase(t *testing.T) {
	//Arrange
	formatter := NewFormatterSvc(validator)
	snakeCaseName := "function_name"

	//Act
	actual := formatter.ToPascalCase(snakeCaseName)

	//Assert
	assert := assert.New(t)

	assert.Equal("FunctionName", actual)
}

func TestToPascalCaseFromKebabCase(t *testing.T) {
	//Arrange
	formatter := NewFormatterSvc(validator)
	kebabCaseName := "function-name"

	//Act
	actual := formatter.ToPascalCase(kebabCaseName)

	//Assert
	assert := assert.New(t)

	assert.Equal("FunctionName", actual)
}

//ToKebabCase casos
func TestToKebabCaseFromKebabCase(t *testing.T) {
	//Arrange
	formatter := NewFormatterSvc(validator)
	camelCaseName := "my-function-name-hello-world"

	//Act
	actual := formatter.ToKebabCase(camelCaseName)

	//Assert
	assert := assert.New(t)

	assert.Equal("my-function-name-hello-world", actual)
}

func TestToKebabCaseFromCamelCase(t *testing.T) {
	//Arrange
	formatter := NewFormatterSvc(validator)
	camelCaseName := "myFunctionNameHelloWorld"

	//Act
	actual := formatter.ToKebabCase(camelCaseName)

	//Assert
	assert := assert.New(t)

	assert.Equal("my-function-name-hello-world", actual)
}

func TestToKebabCaseFromSnakeCase(t *testing.T) {
	//Arrange
	formatter := NewFormatterSvc(validator)
	snakeCaseName := "function_name"

	//Act
	actual := formatter.ToKebabCase(snakeCaseName)

	//Assert
	assert := assert.New(t)

	assert.Equal("function-name", actual)
}

func TestToKebabCaseFromPascalCase(t *testing.T) {
	//Arrange
	formatter := NewFormatterSvc(validator)
	pascalCaseName := "FunctionName"

	//Act
	actual := formatter.ToKebabCase(pascalCaseName)

	//Assert
	assert := assert.New(t)

	assert.Equal("function-name", actual)
}

//ToSnakeCase casos
func TestToSnakeCaseFromSnakeCase(t *testing.T) {
	//Arrange
	formatter := NewFormatterSvc(validator)
	camelCaseName := "function_name"

	//Act
	actual := formatter.ToSnakeCase(camelCaseName)

	//Assert
	assert := assert.New(t)

	assert.Equal("function_name", actual)
}

func TestToSnakeCaseFromCamelCase(t *testing.T) {
	//Arrange
	formatter := NewFormatterSvc(validator)
	camelCaseName := "functionName"

	//Act
	actual := formatter.ToSnakeCase(camelCaseName)

	//Assert
	assert := assert.New(t)

	assert.Equal("function_name", actual)
}

func TestToSnakeCaseFromKebabCase(t *testing.T) {
	//Arrange
	formatter := NewFormatterSvc(validator)
	kebabCaseName := "function-name"

	//Act
	actual := formatter.ToSnakeCase(kebabCaseName)

	//Assert
	assert := assert.New(t)

	assert.Equal("function_name", actual)
}

func TestToSnakeCaseFromPascalCase(t *testing.T) {
	//Arrange
	formatter := NewFormatterSvc(validator)
	pascalCaseName := "FunctionName"

	//Act
	actual := formatter.ToSnakeCase(pascalCaseName)

	//Assert
	assert := assert.New(t)

	assert.Equal("function_name", actual)
}

//Pruebas de casos extremos

func TestToSnakeCaseEmptyInput(t *testing.T) {
	//Arrange
	formatter := NewFormatterSvc(validator)
	pascalCaseName := ""

	//Act
	actual := formatter.ToSnakeCase(pascalCaseName)

	//Assert
	assert := assert.New(t)

	assert.Equal("", actual)
}

func TestToPascalCaseEmptyInput(t *testing.T) {
	//Arrange
	formatter := NewFormatterSvc(validator)
	pascalCaseName := ""

	//Act
	actual := formatter.ToPascalCase(pascalCaseName)

	//Assert
	assert := assert.New(t)

	assert.Equal("", actual)
}

func TestToCamelCaseEmptyInput(t *testing.T) {
	//Arrange
	formatter := NewFormatterSvc(validator)
	pascalCaseName := ""

	//Act
	actual := formatter.ToCamelCase(pascalCaseName)

	//Assert
	assert := assert.New(t)

	assert.Equal("", actual)
}
