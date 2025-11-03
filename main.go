package main

import (
	"callium/core"
)

func main() {

	app := core.New(&core.Config{
		Alias: []string{
			"opensips.org",
			"sip-router.org",
		},
		LogLevel:  "DEBUG",
		UDPWorker: 4,
		TCPWorker: 4,
		Sockets: []string{
			"udp:127.0.0.1:5060",
			"udp:127.0.0.1:5080",
			"udp:127.0.0.1:8090",
		},
	})

	app.OnRequest(func(c *core.Ctx) {
		if c.MaxForwards() <= 10 {
			c.SlSendReply(483, "Too Many Hops")
			return
		}

		if c.Method() == "OPTIONS" {
			c.SlSendReply(200, "Ok")
			return
		}

		if c.Method() == "PUBLISH" || c.Method() == "SUBSCRIBE" {
			c.SlSendReply(503, "Service Unavailable")
			return
		}

		if c.HasToTag() {
			if c.Method() == "ACK" && c.CheckTrans() {
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

		if c.Method() == "CANCEL" {
			if c.CheckTrans() {
				c.TRelay()
			}
			return
		}

		c.CheckTrans()

		if c.Method() != "REGISTER" {
			fd := c.GetVar("$fd")
			c.Infow("fd is", fd)
			c.TRelay()
			return
		}

		c.SlSendReply(404, "Not Found")
	})

	app.OnError(func(c *core.Ctx) {
		c.Errorw("error", c)
	})

	app.OnStarted(func(c *core.Ctx) {
		c.Infow("started")
	})

	app.Start()
}

func relay(c *core.Ctx) {
	c.Infof("enter relay")
}
