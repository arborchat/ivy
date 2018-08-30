package main

import (
    "fmt"
    "net"
    "encoding/json"
)

const (
    proto = "tcp"
    serv = "localhost:7777"
    welcome = "\n--| ~Ivy~ Version: XXX |--\n"
)

func handleit(e error) {
    if e != nil {
        fmt.Println("OH NO: %s", e)
    }
}

type WelcomeMessage struct {
    Type   int
    Root   string
    Recent []string
    Major  int
    Minor  int
}

func main() {
    fmt.Println(welcome)

	var wm WelcomeMessage
    conn, err := net.Dial(proto, serv)
	handleit(err)

	response := json.NewDecoder(conn)

	e := response.Decode(&wm)
	handleit(e)

	fmt.Printf("Type: %d\nRoot Message: %s\nRecent Messages: %v\nServer Version %d.%d\n", wm.Type, wm.Root, wm.Recent, wm.Major, wm.Minor)
}
