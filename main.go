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
	var m, nm Message

	// Server connection
    conn, err := net.Dial(proto, serv)
	handleit(err)

	// JSON decoder
	listener := json.NewDecoder(conn)
	writer   := json.NewEncoder(conn)

	// Get stuff from the server
    fmt.Println("~~ Message from server ~~")
	e := listener.Decode(&m)
	handleit(e)
	fmt.Printf("Type: %d\nRoot Message: %s\nRecent Messages: %v\nServer Version %d.%d\n", m.Type, m.Root, m.Recent, m.Major, m.Minor)

	// Build message for the server
	nm.Type      = 2 // new message
	nm.UUID      = "111-111-111"
	nm.Parent    = m.Root
	nm.Content   = "Hello world!"
	nm.Timestamp = 1
	nm.Username  = "josh"

	// Send stuff to the server
	we := writer.Encode(&nm)
	handleit(we)

    // Get stuff from the server
    fmt.Println("~~ Message from server ~~")
	le := listener.Decode(&m)
	handleit(le)
	fmt.Printf("Type: %d\nParent: %s\nContent: %s\n", m.Type, m.Parent, m.Content)
}
