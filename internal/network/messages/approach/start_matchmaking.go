package approach

import (
	"encoding/binary"
	"io"
	"net"

	"github.com/leikyz/lkz-online-services/internal/models"
	"github.com/leikyz/lkz-online-services/internal/registries"
)

type StartMatchmakingMessage struct {
	ID uint8
}

func NewStartMatchmakingMessage() *StartMatchmakingMessage {
	return &StartMatchmakingMessage{ID: 4}
}

func (m *StartMatchmakingMessage) GetID() uint8 {
	return m.ID
}

func (m *StartMatchmakingMessage) Serialize() ([]byte, error) {
	data := make([]byte, 3)
	data[0] = byte(m.GetMessageSize() >> 8)
	data[1] = byte(m.GetMessageSize() & 0xFF)
	data[2] = m.ID
	return data, nil
}

func (m *StartMatchmakingMessage) Deserialize(reader io.Reader) error {
	return nil
}

func (m *StartMatchmakingMessage) Process(c *models.Client, conn net.Conn) (*models.Client, error) {

	registries.Matchmaking.AddClientToQueue(c)
	return c, nil
}

func (m *StartMatchmakingMessage) GetMessageSize() uint16 {
	return uint16(binary.Size(m) + 2)
}
