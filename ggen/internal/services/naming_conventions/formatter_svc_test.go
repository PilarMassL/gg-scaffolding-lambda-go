package naming_conventions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var validator = NewValidatorSvc()

//ToCamelCase casos
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

func TestToSnakeCaseFromCamelCase(t *testing.T) {
	//Arrange
	formatter := NewFormatterSvc(validator)
	camelCaseName := "FunctionName"

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
