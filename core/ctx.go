package core

import (
	"go.uber.org/zap"
)

type Ctx struct {
	ID int
	Sl
	events map[string]EventListener
	*zap.SugaredLogger
}

func (c *Ctx) Drop() {
	c.Infoln("Drop")
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

func (c *Ctx) GetHeader(h string, index int) int {
	return 70
}

func (c *Ctx) DeleteHeader(h string, index int) int {
	return 70
}

func (c *Ctx) GetVar(k string) string {
	return "abc"
}
func (c *Ctx) SetVar(k, v string) {

}

func (c *Ctx) OnReply(r EventListener) string {
	c.events["OnReply"] = r
	return "reply"
}
func (c *Ctx) OnFailure(r EventListener) string {
	c.events["OnFailure"] = r
	return "failure"
}
