package lobbies

import (
	"encoding/binary"
	"io"
	"net"

	"github.com/leikyz/lkz-online-services/internal/models"
)

type WaitingForSessionMessage struct {
	ID uint8
}

func NewSWaitingForSessionMessage() *WaitingForSessionMessage {
	return &WaitingForSessionMessage{ID: 25}
}

func (m *WaitingForSessionMessage) GetID() uint8 {
	return m.ID
}

func (m *WaitingForSessionMessage) Serialize() ([]byte, error) {
	data := make([]byte, 3)
	data[0] = byte(m.GetMessageSize() >> 8)
	data[1] = byte(m.GetMessageSize() & 0xFF)
	data[2] = m.ID
	return data, nil
}

func (m *WaitingForSessionMessage) Deserialize(reader io.Reader) error {
	return nil
}

func (m *WaitingForSessionMessage) Process(c *models.Client, conn net.Conn) (*models.Client, error) {
	return c, nil
}

func (m *WaitingForSessionMessage) GetMessageSize() uint16 {
	return uint16(binary.Size(m) + 2)
}
