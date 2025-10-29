package core

import "fmt"

type EventListener func(*Ctx)

type Sl struct {
	ID int
}

func (s *Sl) SlSendReply(code int, reason string) {
	fmt.Println("send reply", code, reason)
}

type Ctx struct {
	ID int
	Sl
	events map[string]EventListener
}

func (c *Ctx) Drop() {
	fmt.Println("Drop")
}

func (c *Ctx) GetMethod() string {
	return "GET"
}

func (c *Ctx) HasToTag() bool {
	return false
}

func (c *Ctx) CheckTrans() bool {
	return false
}
func (c *Ctx) TRelay() bool {
	return false
}
func (c *Ctx) LooseRoute() bool {
	return false
}
func (c *Ctx) SendReply(code int, reason string) bool {
	return false
}

func (c *Ctx) MaxForwards() int {
	return 70
}

func (c *Ctx) OnReply(r EventListener) string {
	c.events["OnReply"] = r
	return "reply"
}
func (c *Ctx) OnFailure(r EventListener) string {
	c.events["OnFailure"] = r
	return "failure"
}

type App struct {
	Name    string
	Version string
	events  map[string]EventListener
	Config  Config
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

	app := &App{
		Name:    "Fiber",
		Version: "1.0.0",
		events:  make(map[string]EventListener),
		Config:  c,
	}

	app.rewriteWithDefaults()

	return app
}
