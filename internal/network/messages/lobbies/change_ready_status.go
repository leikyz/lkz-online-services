package lobbies

import (
	"encoding/binary"
	"io"
	"net"

	"github.com/leikyz/lkz-online-services/internal/models"
)

type ChangeReadyStatusMessage struct {
	ID              uint8
	IsReady         bool
	PositionInLobby uint8
}

func NewChangeReadyStatusMessage(isReady bool, positionInLobby uint8) *ChangeReadyStatusMessage {
	return &ChangeReadyStatusMessage{ID: 6, IsReady: isReady, PositionInLobby: positionInLobby}
}

func (m *ChangeReadyStatusMessage) GetID() uint8 {
	return m.ID
}

func (m *ChangeReadyStatusMessage) Serialize() ([]byte, error) {
	data := make([]byte, 4)
	data[0] = byte(m.GetMessageSize() >> 8)
	data[1] = byte(m.GetMessageSize() & 0xFF)
	data[2] = m.ID
	data[3] = m.PositionInLobby
	return data, nil
}

func (m *ChangeReadyStatusMessage) Deserialize(reader io.Reader) error {
	return nil
}

func (m *ChangeReadyStatusMessage) Process(c *models.Client, conn net.Conn) (*models.Client, error) {
	return c, nil
}

func (m *ChangeReadyStatusMessage) GetMessageSize() uint16 {
	return uint16(binary.Size(m) + 2)
}
