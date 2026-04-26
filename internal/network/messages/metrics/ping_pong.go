package metrics

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

type PingPongMessage struct {
	ID         uint8
	SequenceID uint8
}

func NewPingPongMessage() *PingPongMessage {
	return &PingPongMessage{
		ID: 32,
	}
}

func (m *PingPongMessage) GetID() uint8 {
	return m.ID
}

// GetMessageSize returns the total packet size (Header + ID + Payload)
func (m *PingPongMessage) GetMessageSize() uint16 {
	return 4
}

func (m *PingPongMessage) Serialize() ([]byte, error) {
	data := make([]byte, 4)

	// Ensure this matches your Unity client's Endianness
	// BigEndian: [0, 4] | LittleEndian: [4, 0]
	binary.BigEndian.PutUint16(data[0:2], m.GetMessageSize())

	data[2] = m.ID
	data[3] = m.SequenceID
	return data, nil
}

func (m *PingPongMessage) Deserialize(reader io.Reader) error {
	// We only need to read the SequenceID because the
	// ID (32) was already read by your packet dispatcher
	buf := make([]byte, 1)
	_, err := io.ReadFull(reader, buf)
	if err != nil {
		return err
	}
	m.SequenceID = buf[0]
	return nil
}

// Process handles the logic of echoing the ping back to the client
func (m *PingPongMessage) Process(conn net.Conn) error {
	data, err := m.Serialize()
	if err != nil {
		return fmt.Errorf("failed to serialize ping response: %v", err)
	}

	_, err = conn.Write(data)
	return err
}
