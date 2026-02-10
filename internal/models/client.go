package models

import "net"

type Client struct {
	ID       uint32
	Username string
	Level    int
	Conn     net.Conn
	Lobby    *Lobby
	Ready    bool
}
