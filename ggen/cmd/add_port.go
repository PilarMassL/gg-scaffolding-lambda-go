package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// portCmd representa el comando para generar un nuevo puerto a una función
var portCmd = &cobra.Command{
	Use:   "port",
	Short: "Permite generar un Puerto y (opcionalmente) Adapter a una función",
	Long:  `Permite generar un Puerto y (opcionalmente) Adapter a una función`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("port invocado: Pendiente de implementar")

		//Debe validar que:
		//- Está en un proyecto inicializado.
		//- la función indicada exista.
		//- El adapter está dentro de los soportados.
	},
}

func init() {
	addCmd.AddCommand(portCmd)
}
