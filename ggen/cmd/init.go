package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	c "github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/config"
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/generators/project"
)

var projectbase string

// initCmd representa el comando para iniciar un proyecto basado en funciones.
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Permite iniciar un proyecto basado en funciones",
	Long: `Permite iniciar un proyecto basado en funciones.
Normalmente agrupamos en un solo repositorio funciones que mantienen cierta coherencia
y colaboran para ofrecer un servicio que implementa una capacidad del negocio.
Ejemplo: 

ggen init my-project`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init invocado")
		generateProject(projectbase)
		viper.SafeWriteConfig()

	},
}

func init() {
	cobra.OnInitialize(c.InitConfig)
	initCmd.PersistentFlags().StringVarP(&projectbase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	initCmd.MarkPersistentFlagRequired("projectbase")
	viper.BindPFlag("projectbase", initCmd.PersistentFlags().Lookup("projectbase"))
	rootCmd.AddCommand(initCmd)
}

func generateProject(name string) {
	//Construimos todos los objetos que se requieren para la generación.
	projectGenerator := project.NewProjectGenerator(c.BuildBaseGenerator())

	//Pasamos los argumentos a Params.
	params := project.ProjectPromptParams{
		ProjectName: name,
	}

	//Generamos el código fuente.
	_, errGenerating := projectGenerator.Generate(params)
	cobra.CheckErr(errGenerating)

}
