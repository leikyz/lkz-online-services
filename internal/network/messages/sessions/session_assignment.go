package sessions

import (
	"bytes"
	"encoding/binary"
	"io"
)

type SessionAssignmentMessage struct {
	ID    uint8
	Token uint32
}

func NewSessionAssignmentMessage(token uint32) *SessionAssignmentMessage {
	return &SessionAssignmentMessage{
		ID:    27,
		Token: token,
	}
}

func (m *SessionAssignmentMessage) GetID() uint8 {
	return m.ID
}

func (m *SessionAssignmentMessage) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.Grow(int(m.GetMessageSize()))

	binary.Write(buf, binary.BigEndian, uint16(m.GetMessageSize()))

	buf.WriteByte(m.ID)

	binary.Write(buf, binary.BigEndian, m.Token)

	return buf.Bytes(), nil
}

func (m *SessionAssignmentMessage) Deserialize(reader io.Reader) error {
	return nil
}
func (m *SessionAssignmentMessage) GetMessageSize() uint16 {
	return uint16(1 + 2 + 4)
}
