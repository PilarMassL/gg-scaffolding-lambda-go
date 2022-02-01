package config

import (
	"os"

	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/services/generator"
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/services/writer"
	"github.com/spf13/cobra"
)

func BuildBaseGenerator() *generator.GeneratorSvcDefault {
	currentWorkingDirectory, err := os.Getwd()
	cobra.CheckErr(err)
	writer := writer.NewDiskWriter(currentWorkingDirectory)
	return generator.NewGeneratorSvc(writer)
}
