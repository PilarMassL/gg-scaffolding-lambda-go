package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsCamelCaseOnCamelCase(t *testing.T) {
	//Arrange
	helper := NewNamingConventionsHelper()
	camelCaseName := "functionName"
	//Act
	result := helper.IsCamelCase(camelCaseName)

	//Assert
	assert := assert.New(t)
	assert.True(result)
}

func TestIsCamelCaseOnSneakCase(t *testing.T) {
	//Arrange
	helper := NewNamingConventionsHelper()
	snakeCaseName := "function_name"
	//Act
	result := helper.IsCamelCase(snakeCaseName)

	//Assert
	assert := assert.New(t)
	assert.False(result)
}

func TestIsPascalCaseOnPascalCase(t *testing.T) {
	//Arrange
	helper := NewNamingConventionsHelper()
	pascalCaseName := "FunctionName"
	//Act
	result := helper.IsPascalCase(pascalCaseName)

	//Assert
	assert := assert.New(t)
	assert.True(result)
}
func TestIsPascalCaseOnCamelCase(t *testing.T) {
	//Arrange
	helper := NewNamingConventionsHelper()
	camelCaseName := "functionName"
	//Act
	result := helper.IsPascalCase(camelCaseName)

	//Assert
	assert := assert.New(t)
	assert.False(result)
}

func TestIsSnakeCaseOnSnakeCase(t *testing.T) {
	//Arrange
	helper := NewNamingConventionsHelper()
	snakeCaseName := "function_name"
	//Act
	result := helper.IsPascalCase(snakeCaseName)

	//Assert
	assert := assert.New(t)
	assert.True(result)
}

func TestIsSnakeCaseOnPascalCase(t *testing.T) {
	//Arrange
	helper := NewNamingConventionsHelper()
	pascalCaseName := "FunctionName"
	//Act
	result := helper.IsSnakeCase(pascalCaseName)

	//Assert
	assert := assert.New(t)
	assert.False(result)
}

func TestIsKebabCaseOnKebabCase(t *testing.T) {
	//Arrange
	helper := NewNamingConventionsHelper()
	kebabCaseName := "function-name"
	//Act
	result := helper.IsKebabCase(kebabCaseName)

	//Assert
	assert := assert.New(t)
	assert.True(result)
}

func TestIsKebabCaseOnSnakeCase(t *testing.T) {
	//Arrange
	helper := NewNamingConventionsHelper()
	snakeCaseName := "function_name"
	//Act
	result := helper.IsKebabCase(snakeCaseName)

	//Assert
	assert := assert.New(t)
	assert.False(result)
}

//ToCamelCase casos
func TestToCamelCaseFromPascalCase(t *testing.T) {
	//Arrange
	helper := NewNamingConventionsHelper()
	pascalCaseName := "FunctionName"

	//Act
	actual, err := helper.ToCamelCase(pascalCaseName)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal("functionName", actual)
}

func TestToCamelCaseFromSnakeCase(t *testing.T) {
	//Arrange
	helper := NewNamingConventionsHelper()
	snakeCaseName := "function_name"

	//Act
	actual, err := helper.ToCamelCase(snakeCaseName)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal("functionName", actual)
}

func TestToCamelCaseFromKebabCase(t *testing.T) {
	//Arrange
	helper := NewNamingConventionsHelper()
	kebabCaseName := "function-name"

	//Act
	actual, err := helper.ToCamelCase(kebabCaseName)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal("functionName", actual)
}

//ToPascalCase casos

func TestToPascalCaseFromCamelCase(t *testing.T) {
	//Arrange
	helper := NewNamingConventionsHelper()
	camelCaseName := "functionName"

	//Act
	actual, err := helper.ToPascalCase(camelCaseName)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal("FunctionName", actual)
}

func TestToPascalCaseFromSnakeCase(t *testing.T) {
	//Arrange
	helper := NewNamingConventionsHelper()
	snakeCaseName := "function_name"

	//Act
	actual, err := helper.ToPascalCase(snakeCaseName)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal("FunctionName", actual)
}

func TestToPascalCaseFromKebabCase(t *testing.T) {
	//Arrange
	helper := NewNamingConventionsHelper()
	kebabCaseName := "function-name"

	//Act
	actual, err := helper.ToPascalCase(kebabCaseName)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal("FunctionName", actual)
}

//ToKebabCase casos

func TestToKebabCaseFromCamelCase(t *testing.T) {
	//Arrange
	helper := NewNamingConventionsHelper()
	camelCaseName := "FunctionName"

	//Act
	actual, err := helper.ToKebabCase(camelCaseName)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal("function-name", actual)
}

func TestToKebabCaseFromSnakeCase(t *testing.T) {
	//Arrange
	helper := NewNamingConventionsHelper()
	snakeCaseName := "function_name"

	//Act
	actual, err := helper.ToKebabCase(snakeCaseName)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal("function-name", actual)
}

func TestToKebabCaseFromPascalCase(t *testing.T) {
	//Arrange
	helper := NewNamingConventionsHelper()
	pascalCaseName := "FunctionName"

	//Act
	actual, err := helper.ToPascalCase(pascalCaseName)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal("function-name", actual)
}

//ToSnakeCase casos

func TestToSnakeCaseFromCamelCase(t *testing.T) {
	//Arrange
	helper := NewNamingConventionsHelper()
	camelCaseName := "FunctionName"

	//Act
	actual, err := helper.ToSnakeCase(camelCaseName)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal("function_name", actual)
}

func TestToSnakeCaseFromKebabCase(t *testing.T) {
	//Arrange
	helper := NewNamingConventionsHelper()
	kebabCaseName := "function-name"

	//Act
	actual, err := helper.ToKebabCase(kebabCaseName)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal("function_name", actual)
}

func TestToSnakeCaseFromPascalCase(t *testing.T) {
	//Arrange
	helper := NewNamingConventionsHelper()
	pascalCaseName := "FunctionName"

	//Act
	actual, err := helper.ToKebabCase(pascalCaseName)

	//Assert
	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal("function_name", actual)
}
