package metrics

import (
	"encoding/binary"
	"io"
)

type BackendMetricsMessage struct {
	ID           uint8
	LobbyToken   uint32
	PlayerCount  uint16
	SessionToken uint32
}

func NewBackendMetricsMessage(lobbyToken uint32, count uint16, sessionToken uint32) *BackendMetricsMessage {
	return &BackendMetricsMessage{
		ID:           31, // Updated Unique ID
		LobbyToken:   lobbyToken,
		PlayerCount:  count,
		SessionToken: sessionToken,
	}
}

func (m *BackendMetricsMessage) GetID() uint8 {
	return m.ID
}

func (m *BackendMetricsMessage) Serialize() ([]byte, error) {
	size := m.GetMessageSize()
	data := make([]byte, size)

	// 1. Write Header: Size (2) + ID (1)
	binary.BigEndian.PutUint16(data[0:], size)
	data[2] = m.ID

	// 2. Serialize Fixed-Size Fields
	offset := 3
	binary.BigEndian.PutUint32(data[offset:], m.LobbyToken)
	offset += 4

	binary.BigEndian.PutUint16(data[offset:], m.PlayerCount)
	offset += 2

	binary.BigEndian.PutUint32(data[offset:], m.SessionToken)

	return data, nil
}

func (m *BackendMetricsMessage) Deserialize(reader io.Reader) error {
	return nil
}

// Fixed size: 2(size) + 1(ID) + 4(Lobby) + 2(Count) + 4(Session) = 13 bytes
func (m *BackendMetricsMessage) GetMessageSize() uint16 {
	return 13
}
