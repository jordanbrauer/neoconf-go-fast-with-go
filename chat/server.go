package chat

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Server struct {
	connections map[Connection]ID
	joined      chan Connection
	left        chan Connection
	outbox      chan string
}

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

func (server *Server) Listening() int {
	return len(server.connections)
}

func (server *Server) Broadcast(message string) {
	send := func(stream Connection, message string) {
		_, err := stream.Write([]byte(message))

		catch(err)
		log.Printf(
			"%s %s\n",
			blue(stream.ID),
			gray("received a message"),
		)
	}

	for recipient := range server.connections {
		go send(recipient, message)
	}
}

func (server *Server) Leave(stream Connection) {
	log.Printf("%s %s\n", red(stream.ID), gray("has disconnected"))
	stream.Close()
	delete(server.connections, stream)
	log.Printf(gray("Now serving %d client(s)\n"), server.Listening())
}

func (server *Server) Join(stream Connection) {
	server.connections[stream] = stream.ID

	log.Printf("%s %s\n", green(stream.ID), gray("has joined"))
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
