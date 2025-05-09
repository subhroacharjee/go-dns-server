package message

import (
	"bytes"
	"fmt"
	"strings"
)

type RecordType uint16

const (
	A     RecordType = 1  // a host address
	NS    RecordType = 2  // an authoritative name server
	MD    RecordType = 3  // a mail destination (Obsolete - use MX)
	MF    RecordType = 4  // a mail forwarder (Obsolete - use MX)
	CNAME RecordType = 5  // the canonical name for an alias
	SOA   RecordType = 6  // marks the start of a zone of authority
	MB    RecordType = 7  // a mailbox domain name (EXPERIMENTAL)
	MG    RecordType = 8  // a mail group member (EXPERIMENTAL)
	MR    RecordType = 9  // a mail rename domain name (EXPERIMENTAL)
	NULL  RecordType = 10 // a null RR (EXPERIMENTAL)
	WKS   RecordType = 11 // a well known service description
	PTR   RecordType = 12 // a domain name pointer
	HINFO RecordType = 13 // host information
	MINFO RecordType = 14 // mailbox or mail list information
	MX    RecordType = 15 // mail exchange
	TXT   RecordType = 16 // text strings
)

type RecordClass uint16

const (
	IN RecordClass = 1 // the Internet
	CS RecordClass = 2 // the CSNET class (Obsolete - used only for examples in obsolete RFCs)
	CH RecordClass = 3 // the CHAOS class
	HS RecordClass = 4 // Hesiod [Dyer 87]
)

type Question struct {
	Name   string
	_name  []byte
	Type   RecordType
	Class  RecordClass
	_type  []byte
	_class []byte
}

func ParseQuestion(buf []byte) (*Question, []byte) {
	// for now assume it exists
	name, rem, _ := bytes.Cut(buf, []byte{0x00})

	qtype := rem[:2]
	qclass := rem[2:4]

	remBuf := rem[4:]
	fmt.Println(buf)

	// fmt.Println(">>>>>", len(name))

	var nameBuilder strings.Builder
	for i := 0; i < len(name); i++ {
		fmt.Println(i)
		labelSize := int(name[i]) + i
		i++

		for ; i < labelSize && i < len(name); i++ {
			nameBuilder.WriteByte(name[i])
		}

		if i < len(name) {
			nameBuilder.WriteRune('.')
		}

		i--

		// fmt.Println(">>>>>", i)
	}
	name = append(name, 0x00)

	// fmt.Println(name)

	return &Question{
		Name:   nameBuilder.String(),
		_name:  name,
		Type:   RecordType(ByteToUint16(qtype)),
		_type:  qtype,
		Class:  RecordClass(ByteToUint16(qclass)),
		_class: qclass,
	}, remBuf
}

func (q Question) Marshal() []byte {
	l1 := len(q._name)

	buf := make([]byte, l1+4)
	copy(buf[:l1], q._name)
	copy(buf[l1:l1+2], q._type)
	copy(buf[l1+2:l1+4], q._class)
	fmt.Println(buf)
	return buf
}
