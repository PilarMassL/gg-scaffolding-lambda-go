package utils

import (
	"fmt"
	"os"
)

/*
 funciones de ayuda para errores
*/

// CheckErrorFatal recibe un error y detiene la aplicación si no es nil
func CheckErrorFatal(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
