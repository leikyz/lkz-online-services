package messages

import (
	"io"
	"net"

	client "github.com/leikyz/lkz-online-services/internal/registries"
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
	return []byte{m.ID}, nil
}

func (m *CreateClientMessage) Deserialize(reader io.Reader) error {
	return nil
}

func (m *CreateClientMessage) Process(conn net.Conn) error {
	_ = client.Clients.CreateClient("Guest", 0)

	data, _ := m.Serialize()
	_, err := conn.Write(data)

	return err
}
