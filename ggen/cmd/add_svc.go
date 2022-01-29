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
	},
}

func init() {
	addCmd.AddCommand(svcCmd)
}
