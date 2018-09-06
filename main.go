// Chat client for the Arbor protocol.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
)

// Connectivity constants
const (
	proto   = "tcp"
	serv    = "localhost:7777"
	welcome = "\n--| ~Ivy~ Version: XXX |--\n"
)

// Error handler
func handleit(e error) {
	if e != nil {
		fmt.Printf("OH NO: %s", e)
	}
}

// Messages as defined in the Arbor protocol
type Message struct {
	Type      int
	Root      string
	Parent    string
	UUID      string
	Recent    []string
	Content   string
	Username  string
	Timestamp int
	Major     int
	Minor     int
}

func main() {
	// MOTD
	fmt.Println(welcome)

	// Server connection
	conn, err := net.Dial(proto, serv)
	handleit(err)

	// Create channel
	rcvChan := make(chan Message)

	go something(rcvChan, conn)

	// Print stuff from the server sent through the channel
	for incMessage := range rcvChan {
		fmt.Println("~~ Message from server ~~")
		fmt.Printf("Type: %d\nParent: %s\nContent: %s\n", incMessage.Type, incMessage.Parent, incMessage.Content)
	}
}

func something(rcvChan chan<- Message, conn io.Reader) {
	defer close(rcvChan)

	var m Message

	// JSON decoder
	listener := json.NewDecoder(conn)

	// Get root message from the server
	fmt.Println("~~ Message from server ~~")
	e := listener.Decode(&m)
	handleit(e)
	fmt.Printf("Type: %d\nRoot Message: %s\nRecent Messages: %v\nServer Version %d.%d\n", m.Type, m.Root, m.Recent, m.Major, m.Minor)

	// Get stuff from the server and send to channel
	for {
		le := listener.Decode(&m)
		// add handling for the errors?
		if le == io.EOF {
			fmt.Println("Server closed connection...")
			return
		}
		rcvChan <- m
	}
}
