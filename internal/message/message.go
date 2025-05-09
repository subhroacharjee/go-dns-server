package message

type Message struct {
	Header   *Header
	Question *Question
	Answer   Answer
}

func ParseMessage(buf []byte) *Message {
	question, _ := ParseQuestion(buf[12:])
	return &Message{
		Header:   ParseHeader(buf[:12]),
		Question: question,
	}
}

func (m Message) Marshal() []byte {
	buf := make([]byte, 512)
	copy(buf[:12], m.Header.Marshal())
	// fmt.Println("1>>>", m.Header.Marshal())
	// fmt.Println("1>>>", buf)
	question := m.Question.Marshal()

	copy(buf[12:12+len(question)], question)
	// fmt.Println("2>>>", buf)

	return buf
}
