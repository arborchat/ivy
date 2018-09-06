// Chat client for the Arbor protocol.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
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

	go listenFor(rcvChan, conn)
	go sendTo(conn)

	// Print stuff from the server sent through the channel
	for incMessage := range rcvChan {
		printMessage(incMessage)

	}
}

func printMessage(incMessage Message) {
	fmt.Println("~~ Message from server ~~")
	fmt.Printf("Type: %d\nParent: %s\nUUID: %s\nContent: %s\nTime: %v\n",
		incMessage.Type,
		incMessage.Parent,
		incMessage.UUID,
		incMessage.Content,
		incMessage.Timestamp)
}

func listenFor(rcvChan chan<- Message, conn io.Reader) {
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

func sendTo(conn io.Writer) {
	var (
		m     Message
		now   time.Time
		utime int
	)

	readin := bufio.NewReader(os.Stdin)
	writer := json.NewEncoder(conn)

	for {
		time.Sleep(1 * time.Second)

		fmt.Printf("\nReply to: ")
		parent, _ := readin.ReadString('\n')
		fmt.Print("Message: ")
		text, _ := readin.ReadString('\n')

		// Set time for  new message
		now = time.Now()
		utime = int(now.Unix())

		// Set the fields
		m.Type = 2
		m.UUID = "123-123-123"
		m.Parent = strings.TrimSpace(parent)
		m.Content = strings.TrimSpace(text)
		m.Timestamp = utime
		m.Username = "Ivy"

		werr := writer.Encode(&m)
		handleit(werr)
	}
}
