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
		if c.MaxForwards() <= 10 {
			c.SlSendReply(483, "Too Many Hops")
			return
		}

		if c.GetMethod() == "PUBLISH" || c.GetMethod() == "SUBSCRIBE" {
			c.SlSendReply(503, "Service Unavailable")
			return
		}

		if c.HasToTag() {
			if c.GetMethod() == "ACK" && c.CheckTrans() {
				c.TRelay()
				return
			}

			if !c.LooseRoute() {
				// we do record-routing for all our traffic, so we should not
				// receive any sequential requests without Route hdr.
				c.SendReply(404, "Not here")
				return
			}
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

func relay(c *core.Ctx) {

}
