package handlers

import (
	"encoding/binary"
	"io"
	"net"

	"github.com/leikyz/lkz-online-services/internal/network"
	"github.com/leikyz/lkz-online-services/internal/registries"
)

func HandleConnectClient(conn net.Conn, reader io.Reader) error {
	var level int32
	if err := binary.Read(reader, binary.LittleEndian, &level); err != nil {
		return err
	}

	registries.ClientManager.CreateClient("TCP_Player", int(level))

	backendMsg := make([]byte, 5)
	backendMsg[0] = 0x99
	binary.LittleEndian.PutUint32(backendMsg[1:], uint32(level))
	network.Backend.Send(backendMsg)

	resp := []byte{0x01, 0x01}
	_, err := conn.Write(resp)
	return err
}
