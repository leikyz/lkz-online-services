package network

import "net"

type Packet interface {
	Serialize() []byte
	Unserialize(data []byte) error
}

func SendPacket(conn net.Conn, p Packet) error {
	_, err := conn.Write(p.Serialize())
	return err
}
