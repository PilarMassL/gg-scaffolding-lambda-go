package naming_conventions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsCamelCaseOnCamelCase(t *testing.T) {
	//Arrange
	helper := NewValidatorSvc()
	camelCaseName := "functionName"
	//Act
	result := helper.IsCamelCase(camelCaseName)

	//Assert
	assert := assert.New(t)
	assert.True(result)
}

func TestIsCamelCaseOnSneakCase(t *testing.T) {
	//Arrange
	helper := NewValidatorSvc()
	snakeCaseName := "function_name"
	//Act
	result := helper.IsCamelCase(snakeCaseName)

	//Assert
	assert := assert.New(t)
	assert.False(result)
}

func TestIsPascalCaseOnPascalCase(t *testing.T) {
	//Arrange
	helper := NewValidatorSvc()
	pascalCaseName := "FunctionName"
	//Act
	result := helper.IsPascalCase(pascalCaseName)

	//Assert
	assert := assert.New(t)
	assert.True(result)
}
func TestIsPascalCaseOnCamelCase(t *testing.T) {
	//Arrange
	helper := NewValidatorSvc()
	camelCaseName := "functionName"
	//Act
	result := helper.IsPascalCase(camelCaseName)

	//Assert
	assert := assert.New(t)
	assert.False(result)
}

func TestIsSnakeCaseOnSnakeCase(t *testing.T) {
	//Arrange
	helper := NewValidatorSvc()
	snakeCaseName := "function_name"
	//Act
	result := helper.IsSnakeCase(snakeCaseName)

	//Assert
	assert := assert.New(t)
	assert.True(result)
}

func TestIsSnakeCaseOnPascalCase(t *testing.T) {
	//Arrange
	helper := NewValidatorSvc()
	pascalCaseName := "FunctionName"
	//Act
	result := helper.IsSnakeCase(pascalCaseName)

	//Assert
	assert := assert.New(t)
	assert.False(result)
}

func TestIsKebabCaseOnKebabCase(t *testing.T) {
	//Arrange
	helper := NewValidatorSvc()
	kebabCaseName := "function-name"
	//Act
	result := helper.IsKebabCase(kebabCaseName)

	//Assert
	assert := assert.New(t)
	assert.True(result)
}

func TestIsKebabCaseOnSnakeCase(t *testing.T) {
	//Arrange
	helper := NewValidatorSvc()
	snakeCaseName := "function_name"
	//Act
	result := helper.IsKebabCase(snakeCaseName)

	//Assert
	assert := assert.New(t)
	assert.False(result)
}
