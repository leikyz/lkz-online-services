package network

import (
	"io"
	"net"
)

type Message interface {
	GetID() uint8
	Serialize() ([]byte, error)
	Deserialize(reader io.Reader) error
	Process(conn net.Conn) error
}
