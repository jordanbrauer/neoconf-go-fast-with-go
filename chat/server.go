package chat

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// Server represents a set of communication channels for the chat room
type Server struct {
	connections map[Connection]ID
	joined      chan Connection
	left        chan Connection
	outbox      chan string
}

// Listen on the given host and port and begin accepting connections
func (chat *Server) Listen(host string, port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))

	catch(err)
	go func() {
		for {
			connection, err := listener.Accept()

			catch(err)

			chat.joined <- Connection{
				ID:     uuid(),
				Colour: colour(),
				Conn:   connection,
			}
		}
	}()

	for {
		select {
		case user := <-chat.joined:
			chat.Join(user)
		case message := <-chat.outbox:
			chat.Broadcast(message)
		case user := <-chat.left:
			chat.Leave(user)
		}
	}
}

// Listening returns the number of currently connected clients
func (server *Server) Listening() int {
	return len(server.connections)
}

// Broadcast a message to all connected clients
func (server *Server) Broadcast(message string) {
	send := func(stream Connection, message string) {
		_, err := stream.Write([]byte(message))

		catch(err)
		log.Println(blue(stream.ID), gray("received a message"))
	}

	for recipient := range server.connections {
		go send(recipient, message)
	}
}

// Leave the chat room as the given connection
func (server *Server) Leave(stream Connection) {
	catch(stream.Close())
	delete(server.connections, stream)
	log.Println(red(stream.ID), gray("has disconnected"))
	log.Printf(gray("Now serving %d client(s)\n"), server.Listening())
}

// Join the chat room as the given connection
func (server *Server) Join(stream Connection) {
	server.connections[stream] = stream.ID

	log.Println(green(stream.ID), gray("has joined"))
	log.Printf(gray("Now serving %d client(s)\n"), server.Listening())

	go func(stream Connection) {
		id := stream.ID
		reader := bufio.NewReader(stream)
		stream.Write(
			[]byte(fmt.Sprintf("Welcome, %s!\n> ", chalk(id, stream.Colour))),
		)

		for {
			incoming, err := reader.ReadString('\n')

			if nil != err {
				break
			}

			server.outbox <- fmt.Sprintf(
				"\r%s says: %s> ",
				chalk(id, stream.Colour),
				incoming,
			)
		}

		server.left <- stream
	}(stream)
}
