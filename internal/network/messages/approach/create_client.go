package approach

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"

	"github.com/leikyz/lkz-online-services/internal/models"
	"github.com/leikyz/lkz-online-services/internal/network"
	"github.com/leikyz/lkz-online-services/internal/registries"
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
    // Create a new client
    // Use registries.Clients (should be initialized)
    newClient := registries.Clients.CreateClient("Guest", 1, conn)
    
    // Prepare response packet
    data, err := m.Serialize()
    if err != nil {
        return nil, err
    }

    // Send response to the connected client
    // Use 'conn' or 'newClient.Conn' as appropriate
    _, err = conn.Write(data)
    if err != nil {
        return nil, err
    }

    fmt.Printf("Client created successfully: %s\n", newClient.ID)

    // Optionally forward to UDP server if needed
    // Ensure network.Client is not nil before sending
    if network.Client != nil {
        network.Client.Send(data)
    }

    return newClient, nil 
}


func (m *CreateClientMessage) GetMessageSize() uint16 {
	return uint16(binary.Size(m) + 2)
}
