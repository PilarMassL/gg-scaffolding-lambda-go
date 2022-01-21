package port

// MyStoragePort: define un acceso a almacén remoto
type MyStoragePort interface {
	PutFile(key string, content []byte) error
	GetFile(key, version string) ([]byte, error)
}
