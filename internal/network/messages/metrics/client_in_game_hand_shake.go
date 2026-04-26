package metrics

import (
	"encoding/binary"
	"io"
)

type ClientInGameHandShakeMessage struct {
	ID uint8
}

func NewClientInGameHandShakeMessage() *ClientInGameHandShakeMessage {
	return &ClientInGameHandShakeMessage{ID: 22}
}

func (m *ClientInGameHandShakeMessage) GetID() uint8 {
	return m.ID
}

func (m *ClientInGameHandShakeMessage) Serialize() ([]byte, error) {
	data := make([]byte, 3)
	data[0] = byte(m.GetMessageSize() >> 8)
	data[1] = byte(m.GetMessageSize() & 0xFF)
	data[2] = m.ID
	return data, nil
}

func (m *ClientInGameHandShakeMessage) Deserialize(reader io.Reader) error {
	return nil
}

func (m *ClientInGameHandShakeMessage) GetMessageSize() uint16 {
	return uint16(binary.Size(m) + 2)
}
