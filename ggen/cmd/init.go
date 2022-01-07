package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/generators"
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/generators/project"
)

// initCmd representa el comando para iniciar un proyecto basado en funciones.
var initCmd = &cobra.Command{
	Use:   "init [proyect name]",
	Short: "Permite iniciar un proyecto basado en funciones",
	Long: `Permite iniciar un proyecto basado en funciones.
Normalmente agrupamos en un solo repositorio funciones que mantienen cierta coherencia
y colaboran para ofrecer un servicio que implementa una capacidad del negocio.
Ejemplo: 

ggen init my-proyect`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init invocado")

		if len(args) < 1 {
			cobra.CheckErr(fmt.Errorf("init necesita el nombre del proyecto"))
		}

		initProyect(args[0])

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initProyect(name string) {
	wd, err := os.Getwd()
	cobra.CheckErr(err)

	//Construímos todos los objetos que se requieren para la generación.
	writer := generators.NewDiskWriter(wd)
	projectGenerator := project.NewProjectGenerator(writer)

	//Mapeamos los argumentos a Params.
	params := project.ProjectParams{
		ProjectName: name,
	}

	//Generamos el código fuente.
	_, errGenerating := projectGenerator.Generate(params)
	cobra.CheckErr(errGenerating)

}
