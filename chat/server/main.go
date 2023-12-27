package main

import (
	"log"
	"net"

	cs "chat.server/cs"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go cs.Broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go cs.HandleConn(conn)
	}
}
