package main

import (
	"callium/core"
)

func main() {

	app := core.New(&core.Config{
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

			relay(c)
			return
		}

		if c.GetMethod() == "CANCEL" {
			if c.CheckTrans() {
				c.TRelay()
			}
			return
		}

		c.CheckTrans()

		if c.GetMethod() != "REGISTER" {
			c.TRelay()
			return
		}

		c.SlSendReply(404, "Not Found")
	})

	app.Start()
}

func relay(c *core.Ctx) {
	c.Infof("enter relay")
}
