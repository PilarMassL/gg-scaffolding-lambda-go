package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/generators/project"
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/services/generator"
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/services/writer"
)

var projectbase string

// initCmd representa el comando para iniciar un proyecto basado en funciones.
var initCmd = &cobra.Command{
	Use:   "init [project name]",
	Short: "Permite iniciar un proyecto basado en funciones",
	Long: `Permite iniciar un proyecto basado en funciones.
Normalmente agrupamos en un solo repositorio funciones que mantienen cierta coherencia
y colaboran para ofrecer un servicio que implementa una capacidad del negocio.
Ejemplo: 

ggen init my-project`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init invocado")
		initProject(projectbase)

	},
}

func init() {
	cobra.OnInitialize(initConfig)
	initCmd.PersistentFlags().StringVarP(&projectbase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	initCmd.MarkFlagRequired("projectbase")
	viper.BindPFlag("projectbase", initCmd.PersistentFlags().Lookup("projectbase"))
	rootCmd.AddCommand(initCmd)
}

func initProject(name string) {
	currentWorkingDirectory, err := os.Getwd()
	cobra.CheckErr(err)

	//Construimos todos los objetos que se requieren para la generación.
	projectGenerator := buildGenerator(currentWorkingDirectory)

	//Pasamos los argumentos a Params.
	params := project.ProjectPromptParams{
		ProjectName: name,
	}

	//Generamos el código fuente.
	_, errGenerating := projectGenerator.Generate(params)
	cobra.CheckErr(errGenerating)

}

func initConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("ggen")
	viper.SetConfigType("yaml")
	viper.SafeWriteConfig()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("No se puede leer la configuración:", err)
	}
}

func buildGenerator(wd string) *project.ProjectGenerator {
	writer := writer.NewDiskWriter(wd)
	base := generator.NewGeneratorSvc(writer)
	return project.NewProjectGenerator(base)
}
