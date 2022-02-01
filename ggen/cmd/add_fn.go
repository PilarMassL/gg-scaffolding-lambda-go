package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	c "github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/config"
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/generators/fn"
)

var fnName, fnEvent string

// fnCmd representa el comando para generar una nueva función a un proyecto iniciado
var fnCmd = &cobra.Command{
	Use:   "fn",
	Short: "Permite generar una nueva función al proyecto",
	Long:  `Permite generar una nueva función al proyecto. Debe indicar el nombre (name) y tipo de evento (event)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fn invocado")
		c.LoadConfig()
		generateFn(fnName, fnEvent, viper.GetString("projectbase"))
	},
}

func init() {
	cobra.OnInitialize(c.InitConfig)
	addCmd.AddCommand(fnCmd)
	fnCmd.Flags().StringVarP(&fnName, "name", "n", "", "nombre de la función en formato CamelCase, Ej: MyFunctionName")
	fnCmd.Flags().StringVarP(&fnEvent, "event", "e", "generic", "tipo de evento que dispara la función. Por defecto: 'generic'.")
	fnCmd.MarkFlagRequired("name")
}

func generateFn(name, event, baseproject string) {

	//Construimos todos los objetos que se requieren para la generación.
	fnGenerator := fn.NewFnGenerator(c.BuildBaseGenerator())

	//Pasamos los argumentos a Params.
	eventParsed, err := fn.ParseEvent(event)
	cobra.CheckErr(err)

	params := fn.FnPromptParams{
		FnName:        name,
		FnEvent:       eventParsed,
		FnBaseProject: baseproject,
	}

	//Generamos el código fuente.
	_, errGenerating := fnGenerator.Generate(params)
	cobra.CheckErr(errGenerating)

}
