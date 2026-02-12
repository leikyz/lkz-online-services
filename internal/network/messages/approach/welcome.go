package approach

import (
	"encoding/binary"
	"io"
)

type WelcomeMessage struct {
	ID       uint8
	ClientID uint32 // Token
}

func NewWelcomeMessage(clientID uint32) *WelcomeMessage {
	return &WelcomeMessage{ID: 29, ClientID: clientID}
}

func (m *WelcomeMessage) GetID() uint8 {
	return m.ID
}

func (m *WelcomeMessage) Serialize() ([]byte, error) {
	size := m.GetMessageSize()
	data := make([]byte, size)

	data[0] = byte(size >> 8)
	data[1] = byte(size & 0xFF)
	data[2] = m.ID

	binary.BigEndian.PutUint32(data[3:7], m.ClientID)

	return data, nil
}

func (m *WelcomeMessage) Deserialize(reader io.Reader) error {
	return nil
}

func (m *WelcomeMessage) GetMessageSize() uint16 {
	return uint16(binary.Size(m) + 2)
}
