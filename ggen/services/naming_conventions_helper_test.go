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
