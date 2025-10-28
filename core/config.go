package core

type Config struct {
	Alias             []string
	LogLevel          string
	UDPWorker         int
	TCPWorker         int
	Sockets           []string
	TcpMaxConnections int
}
