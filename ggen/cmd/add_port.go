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
	},
}

func init() {
	addCmd.AddCommand(portCmd)
}
