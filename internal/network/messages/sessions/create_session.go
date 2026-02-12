package sessions

import (
	"bytes"
	"encoding/binary"
	"io"
)

type CreateSessionMessage struct {
	ID               uint8
	Token            uint32
	ClientIdsCount   uint8
	ClientAllowedIDs []uint32 // Liste des IDs autorisés à rejoindre
}

func NewCreateSessionMessage(token uint32, clientIDs []uint32) *CreateSessionMessage {
	return &CreateSessionMessage{
		ID:               26,
		Token:            token,
		ClientIdsCount:   uint8(len(clientIDs)),
		ClientAllowedIDs: clientIDs,
	}
}

func (m *CreateSessionMessage) GetID() uint8 {
	return m.ID
}

func (m *CreateSessionMessage) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.Grow(int(m.GetMessageSize()))

	binary.Write(buf, binary.BigEndian, uint16(m.GetMessageSize()))

	buf.WriteByte(m.ID)

	binary.Write(buf, binary.BigEndian, m.Token)

	buf.WriteByte(uint8(len(m.ClientAllowedIDs)))

	for _, clientId := range m.ClientAllowedIDs {
		binary.Write(buf, binary.BigEndian, clientId)
	}

	return buf.Bytes(), nil
}

func (m *CreateSessionMessage) Deserialize(reader io.Reader) error {
	return nil
}
func (m *CreateSessionMessage) GetMessageSize() uint16 {
	fixedValue := uint16(1 + 2 + 4 + 1) // Message ID (1 byte) + message size (2 bytes) + Token (4 bytes) + ClientsIds count

	dynamicValue := uint16(len(m.ClientAllowedIDs) * 4)

	return fixedValue + dynamicValue
}
