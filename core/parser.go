package core

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

func base_parser(buf []byte) (int, error) {
	// Implement your parsing logic here
	return len(buf), nil
}
