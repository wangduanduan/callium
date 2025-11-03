package core

import (
	"strconv"
	"strings"

	"go.uber.org/zap"
)

type App struct {
	events map[string]EventListener
	Config *Config
	*zap.SugaredLogger
	sink0 chan *Packet
}

func (app *App) Start() {

	for _, s := range app.Config.Sockets {
		go app.listen(s)
	}

	for p := range app.sink0 {
		app.Infof("from %s read %v", p.remoteAddr, p.data)
	}
}

func (app *App) listen(socket string) {
	sm := strings.SplitN(socket, ":", 3)

	if len(sm) != 3 {
		app.Fatalw("Invalid socket format", socket)
		return
	}

	port, err := strconv.Atoi(sm[2])
	if err != nil {
		app.Fatalw("Invalid port number", "port", sm[2], "error", err)
		return
	}

	switch sm[0] {
	case "udp":
		s := UdpServer{
			port:                  port,
			ip:                    sm[1],
			SugaredLogger:         app.SugaredLogger,
			maxPacketLength:       app.Config.maxPacketLength,
			minPacketLength:       app.Config.minPacketLength,
			maxReadTimeoutSeconds: app.Config.maxReadTimeoutSeconds,
			output:                app.sink0,
		}
		go s.Serve()
		return
	default:
		app.Fatalw("unsupport socket protocol", sm[0])
		return
	}

}

func (app *App) OnRequest(r EventListener) {
	app.events["OnRequest"] = r
}

func (app *App) OnError(r EventListener) {
	app.events["OnError"] = r
}

func (app *App) OnStarted(r EventListener) {
	app.events["OnStarted"] = r
}

func rewriteWithDefaults(c *Config) {
	if c.Name == "" {
		c.Name = DefaultName
	}

	if c.Version == "" {
		c.Version = DefaultVersion
	}

	if c.LogLevel == "" {
		c.LogLevel = DefaultLogLevel
	}

	if c.UDPWorker <= 0 {
		c.UDPWorker = DefaultUDPWorker
	}

	if c.TCPWorker <= 0 {
		c.TCPWorker = DefaultTCPWorker
	}

	if c.TcpMaxConnections <= 0 {
		c.TcpMaxConnections = DefaultTcpMaxConnections
	}

	if c.Alias == nil {
		c.Alias = []string{}
	}

	if c.Sockets == nil {
		c.Sockets = []string{DefaultListenSocket}
	}

	if c.maxPacketLength <= 0 {
		c.maxPacketLength = DefaultMaxPacketLength
	}

	if c.minPacketLength < 0 {
		c.minPacketLength = DefaultMinPacketLength
	}

	if c.maxReadTimeoutSeconds <= 0 {
		c.maxReadTimeoutSeconds = DefaultMaxReadTimeoutSeconds
	}
}

func New(c *Config) *App {
	rewriteWithDefaults(c)

	InitLogger(c.LogLevel)

	app := &App{
		events:        make(map[string]EventListener),
		Config:        c,
		SugaredLogger: Log,
		sink0:         make(chan *Packet, 100),
	}

	return app
}
