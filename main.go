package main

import (
	"callium/core"
)

func main() {

	app := core.New(core.Config{
		Alias:     []string{"opensips.org", "sip-router.org"},
		LogLevel:  "DEBUG",
		UDPWorker: 4,
		TCPWorker: 4,
		Sockets:   []string{"udp:127.0.0.1:5060"},
	})

	app.OnRequest(func(c *core.Ctx) {
		if c.GetMethod() == "MESSAGE" {
			c.Drop()
			return
		}

		if c.GetMethod() == "INVITE" {
			c.OnReply(func(c *core.Ctx) {
				c.Drop()
			})

			c.OnFailure(func(c *core.Ctx) {
				c.Drop()
			})
		}

		c.SlSendReply(404, "Not Found")
	})

	app.Start()
}
