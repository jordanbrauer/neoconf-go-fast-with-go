package chat

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type ID = string

// New up a chat server
func New() *Server {
	return &Server{
		joined:      make(chan Connection),
		left:        make(chan Connection),
		connections: make(map[Connection]string),
		outbox:      make(chan string),
	}
}

func catch(caught error) {
	if nil != caught {
		log.Panic(caught)
	}
}

func uuid() string {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)

	if nil != err {
		log.Panic(err)
	}

	return fmt.Sprintf(
		"%x%x%x%x%x",
		bytes[0:4],
		bytes[4:6],
		bytes[6:8],
		bytes[8:10],
		bytes[10:],
	)[:8]
}

func colour() uint8 {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	return uint8(21 + random.Intn(210))
}

func chalk(text string, colour uint8) string {
	return fmt.Sprintf("\033[38;5;%dm%s\033[0m", colour, text)
}

func red(text string) string {
	return chalk(text, 1)
}

func green(text string) string {
	return chalk(text, 2)
}

func blue(text string) string {
	return chalk(text, 45)
}

func gray(text string) string {
	return chalk(text, 243)
}
