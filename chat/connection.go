package chat

import "net"

// Connection represents a single client connected to the server
type Connection struct {
	net.Conn

	ID     ID
	Colour uint8
}
