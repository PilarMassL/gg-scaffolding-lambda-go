package writer

import (
	"github.com/PilarMassL/gg-scaffolding-lambda-go/ggen/internal/models"
)

// Writer define un objeto con la capacidad de guardar archivos de código fuente generados.
type Writer interface {
	Save([]models.SrcTpl) ([]models.SrcFile, error)
}
