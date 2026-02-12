package lobbies

import (
	"encoding/binary"
	"io"
)

type JoinLobbyMessage struct {
	ID              uint8
	PositionInLobby uint8
}

func NewJoinLobbyMessage(positionInLobby uint8) *JoinLobbyMessage {
	return &JoinLobbyMessage{ID: 8, PositionInLobby: positionInLobby}
}

func (m *JoinLobbyMessage) GetID() uint8 {
	return m.ID
}

func (m *JoinLobbyMessage) Serialize() ([]byte, error) {
	data := make([]byte, 4)
	data[0] = byte(m.GetMessageSize() >> 8)
	data[1] = byte(m.GetMessageSize() & 0xFF)
	data[2] = m.ID
	data[3] = m.PositionInLobby
	return data, nil
}

func (m *JoinLobbyMessage) Deserialize(reader io.Reader) error {
	return nil
}

func (m *JoinLobbyMessage) GetMessageSize() uint16 {
	return uint16(binary.Size(m) + 2)
}
