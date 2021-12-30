package port

type MyIpClientPort interface {
	GetIp() (string, error)
}
