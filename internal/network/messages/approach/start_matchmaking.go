package approach

import (
	"encoding/binary"
	"io"
	"net"
"fmt"
	"github.com/leikyz/lkz-online-services/internal/models"
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
	data, _ := m.Serialize()
	_, err := conn.Write(data)

	return c, err}

func (m *StartMatchmakingMessage) GetMessageSize() uint16 {
	return uint16(binary.Size(m) + 2)
}
