package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// svcCmd representa el comando para generar un nuevo servicio de dominio a una función
var svcCmd = &cobra.Command{
	Use:   "svc",
	Short: "Permite generar un servicio a la función",
	Long:  `Permite generar un servicio a la función`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("svc invocado, Pendiente de implementar")
		//Debe validar que:
		//- Está en un proyecto inicializado.
		//- la función indicada exista.

	},
}

func init() {
	addCmd.AddCommand(svcCmd)
}
