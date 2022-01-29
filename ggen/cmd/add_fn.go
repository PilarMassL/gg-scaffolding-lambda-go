package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// fnCmd representa el comando para generar una nueva función a un proyecto iniciado
var fnCmd = &cobra.Command{
	Use:   "fn",
	Short: "Permite generar una nueva función al proyecto",
	Long:  `Permite generar una nueva función al proyecto`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fn invocado")
	},
}

func init() {
	addCmd.AddCommand(fnCmd)
}
