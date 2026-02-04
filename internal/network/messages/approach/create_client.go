package approach

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"

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

func (m *CreateClientMessage) Process(conn net.Conn) error {

	newClient := registries.Clients.CreateClient("Guest", 0)
	data, _ := m.Serialize()
	_, err := conn.Write(data)

	if err == nil {
		fmt.Printf("Client created successfuly: %s\n", newClient.ID)
	}

	network.Client.Send(data)

	return err
}

func (m *CreateClientMessage) GetMessageSize() uint16 {
	return uint16(binary.Size(m) + 2)
}
