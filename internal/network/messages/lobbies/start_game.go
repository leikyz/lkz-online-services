package lobbies

import (
	"encoding/binary"
	"io"
	"net"

	"github.com/leikyz/lkz-online-services/internal/models"
)

type StartGameMessage struct {
	ID uint8
}

func NewStartMatchmakingMessage() *StartGameMessage {
	return &StartGameMessage{ID: 5}
}

func (m *StartGameMessage) GetID() uint8 {
	return m.ID
}

func (m *StartGameMessage) Serialize() ([]byte, error) {
	data := make([]byte, 3)
	data[0] = byte(m.GetMessageSize() >> 8)
	data[1] = byte(m.GetMessageSize() & 0xFF)
	data[2] = m.ID
	return data, nil
}

func (m *StartGameMessage) Deserialize(reader io.Reader) error {
	return nil
}

func (m *StartGameMessage) Process(c *models.Client, conn net.Conn) (*models.Client, error) {
	return c, nil
}

func (m *StartGameMessage) GetMessageSize() uint16 {
	return uint16(binary.Size(m) + 2)
}
