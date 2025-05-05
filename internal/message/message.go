package message

type Message struct {
	Header   *Header
	Question Question
	Answer   Answer
}

func ParseMessage(buf []byte) *Message {
	return &Message{
		Header: ParseHeader(buf[:12]),
	}
}

func (m Message) Marshal() []byte {
	buf := make([]byte, 512)
	copy(buf[:12], m.Header.Marshal())
	return buf
}
