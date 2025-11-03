package core

import (
	"bytes"
	"errors"
)

const (
	CRLF = "\r\n"
)

type Message struct {
	Method  string
	Version string
	URI     string
	Status  int // 0 for request, >0 for reply
	Reason  string

	CallID string

	Headers []string

	Body []byte
	raw  []byte
}

func base_parser0(buf []byte) ([][]byte, []byte, error) {
	lines := bytes.Split(buf, []byte(CRLF+CRLF))

	if len(lines) != 2 {
		return nil, nil, errors.New("no body delimiter")
	}

	body := lines[1]

	headers := bytes.Split(lines[0], []byte(CRLF))

	if len(headers) < 1 {
		return nil, nil, errors.New("no line delimiter")
	}

	return headers, body, nil

}
