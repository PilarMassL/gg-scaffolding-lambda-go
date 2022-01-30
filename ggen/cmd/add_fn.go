package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/generators/fn"
)

var fnName, fnEvent string

// fnCmd representa el comando para generar una nueva función a un proyecto iniciado
var fnCmd = &cobra.Command{
	Use:   "fn",
	Short: "Permite generar una nueva función al proyecto",
	Long:  `Permite generar una nueva función al proyecto`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fn invocado")
		generateFn(fnName, fnEvent)
	},
}

func init() {
	addCmd.AddCommand(fnCmd)
	fnCmd.Flags().StringVarP(&fnName, "name", "n", "", "function name")
	fnCmd.Flags().StringVarP(&fnEvent, "event", "e", "generic", "event that triggers the function")
	fnCmd.MarkFlagRequired("name")
}

func generateFn(name, event string) {

	//Construimos todos los objetos que se requieren para la generación.
	fnGenerator := fn.NewFnGenerator(BuildBaseGenerator())

	//Pasamos los argumentos a Params.
	eventParsed, err := fn.ParseEvent(event)
	cobra.CheckErr(err)
	params := fn.FnPromptParams{
		FnName:  name,
		FnEvent: eventParsed,
	}

	//Generamos el código fuente.
	_, errGenerating := fnGenerator.Generate(params)
	cobra.CheckErr(errGenerating)

}
