package cmd

import (
	"errors"
	"fmt"
	"os"
)

// PromptInteractively: recibe un puntero a una estructura de datos anotada adecuadamente y realiza los prompts
// necesarios al usuario para obtener los valores de cada uno de los campos.
func PromptInteractively(dto *struct{}) error {
	return errors.New("not implemented")
}

type Prompt interface {
	Run() (int, string, error)
}

// TODO: resolver el problema de obtener los metadatos de la estructura y construir un slice Prompts.
// Volver el método privado
func StartPrompting(prompts []Prompt) {
	/*
		prompt := promptui.Select{
			Label: "Select Day",
			Items: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
				"Saturday", "Sunday"},
		} */

	for _, p := range prompts {
		_, result, err := p.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("You choose %q\n", result)
	}
}
