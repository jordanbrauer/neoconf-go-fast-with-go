package chat

import "net"

type Connection struct {
	net.Conn

	ID     ID
	Colour uint8
}
