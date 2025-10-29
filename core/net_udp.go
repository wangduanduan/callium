package core

import (
	"net"
	"time"

	"go.uber.org/zap"
)

type UdpServer struct {
	conn *net.UDPConn
	port int
	*zap.SugaredLogger
	maxPacketLength       int
	minPacketLength       int
	maxReadTimeoutSeconds int
	output                chan *Packet
}

type Packet struct {
	data       []byte
	remoteAddr *net.UDPAddr
}

func (u *UdpServer) Serve() {
	var err error
	u.conn, err = net.ListenUDP("udp", &net.UDPAddr{Port: u.port})

	if err != nil {
		u.Fatalf("Udp Service listen report udp fail:%v", err)
	}

	u.Infof("create udp success, listen udp port %v", u.port)

	defer u.conn.Close()
	var data = make([]byte, u.maxPacketLength)

	for {
		u.conn.SetDeadline(time.Now().Add(time.Duration(u.maxReadTimeoutSeconds) * time.Second))
		n, remoteAddr, err := u.conn.ReadFromUDP(data)

		if err != nil {
			if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
				continue
			} else {
				u.Errorf("read udp error: %v, from %v", err, remoteAddr)
			}
		}

		if n < u.minPacketLength {
			u.Warnf("less then minPacketLength: %d, received length: %d, from: %v", u.minPacketLength, n, remoteAddr)
			continue
		}

		raw := make([]byte, n)

		copy(raw, data[:n])

		pkt := &Packet{
			data:       raw,
			remoteAddr: remoteAddr,
		}

		u.output <- pkt
	}
}
