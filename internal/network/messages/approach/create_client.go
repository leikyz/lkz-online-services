package approach

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"

	"github.com/leikyz/lkz-online-services/internal/models"
)

type CreateClientMessage struct {
	ID uint8
}

func NewCreateClientMessage() *CreateClientMessage {
	return &CreateClientMessage{ID: 1}
}

func (m *CreateClientMessage) GetID() uint8 {
	return m.ID
}

func (m *CreateClientMessage) Serialize() ([]byte, error) {
	data := make([]byte, 3)
	data[0] = byte(m.GetMessageSize() >> 8)
	data[1] = byte(m.GetMessageSize() & 0xFF)
	data[2] = m.ID
	return data, nil
}

func (m *CreateClientMessage) Deserialize(reader io.Reader) error {
	return nil
}

func (m *CreateClientMessage) Process(c *models.Client, conn net.Conn) (*models.Client, error) {
    // Create a new client object
    newClient := &models.Client{
        ID:       fmt.Sprintf("Guest-%d", 1),
        Username: "Guest",
        Level:    1,
        Conn:     conn,
    }
    
    // Prepare response packet
    data, err := m.Serialize()
    if err != nil {
        return nil, err
    }

    // Send response to the connected client
    _, err = conn.Write(data)
    if err != nil {
        return nil, err
    }

    fmt.Printf("Client created successfully: %s\n", newClient.ID)

    return newClient, nil 
}


func (m *CreateClientMessage) GetMessageSize() uint16 {
	return uint16(binary.Size(m) + 2)
}
