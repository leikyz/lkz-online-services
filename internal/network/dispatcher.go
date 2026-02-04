package network

import (
	"fmt"
	"io"
	"net"
)

type Message interface {
	GetID() uint8
	Serialize() ([]byte, error)
	Deserialize(reader io.Reader) error
	Process(conn net.Conn) error
}

var registry = make(map[uint8]func() Message)

func RegisterMessage(id uint8, factory func() Message) {
	registry[id] = factory
}

// HandleMessage processes incoming messages from a connection
func HandleMessage(conn net.Conn) {
	defer conn.Close() // We ensure the connection is closed when done

	for {
		idBuf := make([]byte, 3)

		if _, err := io.ReadFull(conn, idBuf); err != nil {
			break
		}

		msgID := idBuf[2] // 0 and 1 are for length, 2 is for message ID

		factory, ok := registry[msgID]

		if !ok {
			fmt.Printf("Message ID inconnu : %d\n", msgID)
			continue
		}

		// Create a new message instance
		msg := factory()

		// Deserialize the message
		if err := msg.Deserialize(conn); err != nil {
			break
		}

		// Process the message
		if err := msg.Process(conn); err != nil {
			break
		}
	}
}
