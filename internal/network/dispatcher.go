package network

import (
	"fmt"
	"io"
	"net"

	"github.com/leikyz/lkz-online-services/internal/models"
)

type Message interface {
	GetID() uint8
	Serialize() ([]byte, error)
	Deserialize(reader io.Reader) error
	GetMessageSize() uint16
}

type HandlerFunc func(msg Message, c *models.Client, conn net.Conn) (*models.Client, error)

var MessageRegistry = make(map[uint8]func() Message)
var HandlerRegistry = make(map[uint8]HandlerFunc)

func Register(id uint8, factory func() Message, handler HandlerFunc) {
	MessageRegistry[id] = factory
	HandlerRegistry[id] = handler
}

func Bind[T Message](handler func(T, *models.Client, net.Conn) (*models.Client, error)) HandlerFunc {
	return func(m Message, c *models.Client, conn net.Conn) (*models.Client, error) {
		specificMsg, ok := m.(T)
		if !ok {
			return nil, fmt.Errorf("expected %T but got %T", new(T), m)
		}
		return handler(specificMsg, c, conn)
	}
}

func HandleMessage(conn net.Conn) {
	defer conn.Close()

	var sessionClient *models.Client

	for {
		// Read header: 2 bytes for size + 1 byte for ID
		headerBuf := make([]byte, 3)
		if _, err := io.ReadFull(conn, headerBuf); err != nil {
			// Connection closed or read error
			break
		}
		msgID := headerBuf[2]
		fmt.Printf("Received Message ID: %d\n", msgID)

		// 1. Get Message Factory
		messageFactory, ok := MessageRegistry[msgID]
		if !ok {
			fmt.Printf("Unknown Message ID: %d\n", msgID)
			continue
		}
		msg := messageFactory()

		// Deserialize payload from connection
		if err := msg.Deserialize(conn); err != nil {
			fmt.Printf("Deserialization error for msg %d: %v\n", msgID, err)
			break
		}

		// Get Logic Handler
		handler, ok := HandlerRegistry[msgID]
		if !ok {
			fmt.Printf("No handler registered for message ID %d\n", msgID)
			continue
		}

		// Special case: Initial Connection / Client Creation (ID 1)
		if msgID == 1 {
			if sessionClient != nil {
				fmt.Println("Client already identified for this connection. Ignoring.")
				continue
			}

			// Pass nil because sessionClient is not yet created
			client, err := handler(msg, nil, conn)
			if err != nil {
				fmt.Printf("Logic error during identification: %v\n", err)
				continue
			}
			sessionClient = client
			continue
		}

		// Security check: Ensure client is identified for any other messages
		if sessionClient == nil {
			fmt.Printf("Action denied for msg %d: Client not identified\n", msgID)
			continue
		}

		// Execute Logic using the Handler (Decoupled from the message struct)
		_, err := handler(msg, sessionClient, conn)
		if err != nil {
			fmt.Printf("Logic error for msg %d: %v\n", msgID, err)
			continue
		}
	}

	// Clean up on disconnect
	//if sessionClient != nil {
	//fmt.Printf("Client disconnected: %s\n", sessionClient.ID)
	//registries.Clients.Remove(sessionClient.ID)
	//}
}
