package server

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/dns-server-starter-go/internal/message"
)

type DNSServer struct {
	addr *net.UDPAddr
}

func NewDnsServer(addr *net.UDPAddr) *DNSServer {
	return &DNSServer{
		addr,
	}
}

func (s *DNSServer) Serve() error {
	udpConn, err := net.ListenUDP("udp", s.addr)
	if err != nil {
		fmt.Println("Failed to bind to address:", err)
		return err
	}
	defer udpConn.Close()

	buf := make([]byte, 512)

	for {

		size, source, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error receiving data:", err)
			return err
		}

		recievedPacket := message.ParseMessage(buf[:size])

		qdcount := 0
		if recievedPacket.Question != nil {
			qdcount = 1
		}

		header := &message.Header{
			ID:      recievedPacket.Header.ID,
			Flag:    message.NewFlag([]byte{0x00, 0x00}),
			QDCount: uint16(qdcount),
			ANCount: 0,
			NSCount: 0,
			ARCount: 0,
		}

		header.Flag.SetQR(true)

		msg := message.Message{
			Header:   header,
			Question: recievedPacket.Question,
		}

		if _, err := udpConn.WriteToUDP(msg.Marshal(), source); err != nil {
			fmt.Println("Error sending message")
		}
	}
}
