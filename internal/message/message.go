package message

type Message struct {
	Header   *Header
	Question *Question
	Answer   *Answer
}

func ParseMessage(buf []byte) *Message {
	question := ParseQuestion(buf[12:])
	return &Message{
		Header:   ParseHeader(buf[:12]),
		Question: question,
	}
}

func (m Message) Marshal() []byte {
	buf := make([]byte, 0)
	buf = append(buf, m.Header.Marshal()...)
	// fmt.Println("1>>>", m.Header.Marshal())
	// fmt.Println("1>>>", buf)
	// question := m.Question.Marshal()

	buf = append(buf, m.Question.Marshal()...)
	// fmt.Println("2>>>", buf)
	//
	// answer := m.Answer.Marsal()
	// copy(buf[sz:sz+len(answer)], answer)
	//
	buf = append(buf, m.Answer.Marsal()...)
	// fmt.Println(buf)

	return buf
}
