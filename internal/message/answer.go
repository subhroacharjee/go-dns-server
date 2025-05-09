package message

import (
	"encoding/binary"
	"fmt"
)

type Answer struct {
	q      *Question
	TTL    []byte
	Length []byte
	Data   []byte
}

func NewAnswer(q *Question) *Answer {
	ttl := make([]byte, 4)
	length := make([]byte, 2)

	binary.BigEndian.PutUint32(ttl, uint32(60))
	binary.BigEndian.PutUint16(length, uint16(4))

	return &Answer{
		q:      q,
		TTL:    ttl,
		Length: length,
		Data:   []byte{0x08, 0x08, 0x08, 0x08},
	}
}

func (a Answer) Marsal() []byte {
	buf := a.q.Marshal()
	buf = append(buf, a.TTL...)
	buf = append(buf, a.Length...)
	buf = append(buf, a.Data...)

	fmt.Println(buf)
	return buf
}
