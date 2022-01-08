package port

import (
	"os"
)

// MyStoragePort: define un acceso a almacén remoto
type MyStoragePort interface {
	PutFile(key string, file *os.File, contentType, perm string) error
	GetFile(key, version string) (*os.File, error)
}
