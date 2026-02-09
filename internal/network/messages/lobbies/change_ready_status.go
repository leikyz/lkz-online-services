package lobbies

import (
	"bytes"
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
	buf := new(bytes.Buffer)
	buf.Grow(5)

	binary.Write(buf, binary.BigEndian, uint16(m.GetMessageSize()))
	buf.WriteByte(m.ID)

	var isReady uint8

	if m.IsReady {
		isReady = 1
	} else {
		isReady = 0
	}

	buf.WriteByte(isReady)
	buf.WriteByte(m.PositionInLobby)

	return buf.Bytes(), nil
}

func (m *ChangeReadyStatusMessage) Deserialize(reader io.Reader) error {
	return nil
}

func (m *ChangeReadyStatusMessage) Process(c *models.Client, conn net.Conn) (*models.Client, error) {

	if c.Lobby == nil {
		return c, nil
	}

	c.Ready = !c.Ready

	var isEveryonReady bool = true

	for i, client := range c.Lobby.Clients {

		if !client.Ready {
			isEveryonReady = false
		}

		message := NewChangeReadyStatusMessage(c.Ready, uint8(i))
		data, _ := message.Serialize()
		client.Conn.Write(data)
	}

	// If everyon is ready, we send StartGame message to all
	if isEveryonReady {
		message := NewSWaitingForSessionMessage()
		data, _ := message.Serialize()

		for _, client := range c.Lobby.Clients {
			client.Conn.Write(data)
		}
	}

	return c, nil
}

func (m *ChangeReadyStatusMessage) GetMessageSize() uint16 {
	return uint16(binary.Size(m) + 2)
}
