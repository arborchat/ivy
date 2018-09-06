// Chat client for the Arbor protocol.
package main

import (
    "fmt"
    "net"
    "encoding/json"
)

// Connectivity constants
const (
    proto = "tcp"
    serv = "localhost:7777"
    welcome = "\n--| ~Ivy~ Version: XXX |--\n"
)

// Error handler
func handleit(e error) {
    if e != nil {
        fmt.Println("OH NO: %s", e)
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

	// Declare messages
	var m Message

	// Server connection
    conn, err := net.Dial(proto, serv)
	handleit(err)


	// JSON decoder
	listener := json.NewDecoder(conn)


	// Get root message from the server
    fmt.Println("~~ Message from server ~~")
	e := listener.Decode(&m)
	handleit(e)
	fmt.Printf("Type: %d\nRoot Message: %s\nRecent Messages: %v\nServer Version %d.%d\n", m.Type, m.Root, m.Recent, m.Major, m.Minor)

	// Create channel
	rcvChan := make(chan Message)

	go func() {
        // Get stuff from the server and send to channel
        for {
        	le := listener.Decode(&m)
        	handleit(le)
        	rcvChan <- m
        }
	}()

	// Print stuff from the server sent through the channel
	for {
    	incMessage := <-rcvChan
        fmt.Println("~~ Message from server ~~")
    	fmt.Printf("Type: %d\nParent: %s\nContent: %s\n", incMessage.Type, incMessage.Parent, incMessage.Content)
	}
}
