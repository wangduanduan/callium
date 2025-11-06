package core

import (
	"bytes"
	"errors"
	"strconv"
)

const (
	CRLF        = "\r\n"
	SIP_VERSION = "SIP/2.0"
)

type HDR struct {
	T           int // type
	Value       string
	raw         []byte
	dirty       bool
	fullParesed bool
}

type Message struct {
	Method  int
	Version string
	URI     string
	Status  int // 0 for request, >0 for reply
	Reason  string
	CallID  string
	Headers map[int][]*HDR

	firstLine []byte
	headers   [][]byte
	body      []byte
	raw       []byte
}

func (m *Message) parseShadowHeaders() {
	for _, h := range m.headers {
		// TODO: parse shadow headers
	}
}

func (m *Message) BaseParse() error {
	lines := bytes.Split(m.raw, []byte(CRLF+CRLF))

	if len(lines) != 2 {
		return errors.New("no body delimiter")
	}

	m.body = lines[1]

	headers := bytes.Split(lines[0], []byte(CRLF))

	if len(headers) < 4 {
		return errors.New("no line delimiter")
	}

	m.firstLine = headers[0]
	m.headers = headers[1:]

	if bytes.HasPrefix(m.firstLine, []byte(SIP_VERSION)) {
		return m.parseStatusLine()
	}

	return m.parseRequestLine()
}

// REGISTER sips:ss2.biloxi.example.com SIP/2.0
func (m *Message) parseRequestLine() error {
	parts := bytes.SplitN(m.firstLine, []byte(" "), 3)
	if len(parts) != 3 {
		return errors.New("malformed request line")
	}

	method := string(parts[0])
	uri := string(parts[1])
	version := string(parts[2])

	if version != SIP_VERSION {
		return errors.New("unsupported SIP version")
	}

	v, ok := smName2Type[method]

	if !ok {
		return errors.New("unknown SIP method")
	}

	m.Method = v
	m.URI = uri
	m.Version = version

	return nil
}

// SIP/2.0 401 Unauthorized
func (m *Message) parseStatusLine() error {
	parts := bytes.SplitN(m.firstLine, []byte(" "), 3)
	if len(parts) < 3 {
		return errors.New("malformed status line")
	}
	m.Version = string(parts[0])
	m.Reason = string(parts[2])

	statusCodeBytes := string(parts[1])

	if len(statusCodeBytes) != 3 {
		return errors.New("malformed status code")
	}

	code, err := strconv.Atoi(statusCodeBytes)

	if err != nil {
		return err
	}

	if code < 100 || code > 699 {
		return errors.New("status code out of range")
	}

	m.Status = code

	return nil
}
