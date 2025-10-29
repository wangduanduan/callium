package core

import (
	"fmt"

	"go.uber.org/zap"
)

type App struct {
	Name    string
	Version string
	events  map[string]EventListener
	Config  Config
	*zap.SugaredLogger
}

func (app *App) Start() {
	fmt.Println("Start")
}

func (app *App) OnRequest(r EventListener) {
	fmt.Println("Stop")
	app.events["OnRequest"] = r
}

func (app *App) rewriteWithDefaults() {
	fmt.Println("rewriteWithDefaults")
}

func New(c Config) *App {
	fmt.Println("New")

	InitLogger(c.LogLevel)

	app := &App{
		Name:          "Callium",
		Version:       "1.0.0",
		events:        make(map[string]EventListener),
		Config:        c,
		SugaredLogger: Log,
	}

	app.rewriteWithDefaults()

	return app
}
