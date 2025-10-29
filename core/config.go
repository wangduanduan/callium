package core

type Config struct {
	Name              string
	Version           string
	Alias             []string
	LogLevel          string
	UDPWorker         int
	TCPWorker         int
	Sockets           []string
	TcpMaxConnections int

	maxPacketLength       int
	minPacketLength       int
	maxReadTimeoutSeconds int
}

const (
	DefaultName                  = "Callium"
	DefaultVersion               = "unknown"
	DefaultLogLevel              = "debug"
	DefaultUDPWorker             = 4
	DefaultTCPWorker             = 4
	DefaultTcpMaxConnections     = 10000
	DefaultMaxPacketLength       = 2048
	DefaultMinPacketLength       = 105
	DefaultMaxReadTimeoutSeconds = 5
	DefaultListenSocket          = "udp:127.0.0.1:8080"
)
