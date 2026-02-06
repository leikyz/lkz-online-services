package models

import "net"

type Client struct {
	ID       string
	Username string
	Level    int
	Conn	 net.Conn
}
