package message

import (
	"encoding/binary"
	"fmt"
	"strings"
)

func ByteToUint16(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}

func EncodeName(s string) []byte {
	labels := strings.Split(s, ".")
	buf := make([]byte, 0)

	for _, label := range labels {
		size := len(label)
		updatedLabel := []byte(fmt.Sprintf("%d%s", size, label))
		buf = append(buf, updatedLabel...)
	}
	buf = append(buf, 0x00)
	return buf
}
