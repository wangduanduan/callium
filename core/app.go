package core

import (
	"fmt"

	"go.uber.org/zap"
)

type App struct {
	events map[string]EventListener
	Config *Config
	*zap.SugaredLogger
}

func (app *App) Start() {
	fmt.Println("Start")
}

func (app *App) OnRequest(r EventListener) {
	fmt.Println("Stop")
	app.events["OnRequest"] = r
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
	}

	return app
}
