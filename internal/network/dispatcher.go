package network

import (
	"fmt"
	"io"
	"net"
)

var registry = make(map[uint8]func() Message)

func RegisterMessage(id uint8, factory func() Message) {
	registry[id] = factory
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		idBuf := make([]byte, 1)
		if _, err := io.ReadFull(conn, idBuf); err != nil {
			break
		}
		msgID := idBuf[0]

		factory, ok := registry[msgID]
		if !ok {
			fmt.Printf("Message ID inconnu : %d\n", msgID)
			continue
		}

		msg := factory()

		if err := msg.Deserialize(conn); err != nil {
			break
		}

		if err := msg.Process(conn); err != nil {
			break
		}
	}
}
