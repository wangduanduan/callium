package core

import "fmt"

type Sl struct {
	ID int
}

func (s *Sl) SlSendReply(code int, reason string) {
	fmt.Println("send reply", code, reason)
}
