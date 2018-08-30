package main

import (
    "fmt"
    "net"
)

func main() {
    fmt.Println("AYEEE LOL")
    conn, err := net.Dial("tcp", "localhost:7777")

    if err != nil {
        fmt.Println("OH NO: %s", err)
    }

	bs := make([]byte, 256)

    iread, _ := conn.Read(bs)

	fmt.Printf("I read %d bytes. Here's the message:\n %s", iread, string(bs))
}
