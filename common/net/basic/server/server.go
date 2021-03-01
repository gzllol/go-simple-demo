package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func handleConnection(conn net.Conn) {
	fmt.Printf("new conn on %v - %v", conn.LocalAddr().String(), conn.RemoteAddr().String())
	fmt.Fprintf(conn, "hello from server\n")
	conn.Close()
}

func main() {
	log.Fatal(http.ListenAndServe(":8081", http.FileServer(http.Dir("/usr/share/doc"))))

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handleConnection(conn)
	}
}
