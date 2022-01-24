package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd representa el comando para agregar un nuevo elemento de arquitectura.
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Permite agregar un nuevo elemento de arquitectura",
	Long: `Permite agregar un nuevo elemento de arquitectura:

fn:   agrega una nueva función.
port: agrega un nuevo puerto/adaptador a una función.
svc:  agrega un nuevo servicio de dominio a la función.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add invocado")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
