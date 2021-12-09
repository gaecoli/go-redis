package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func ListenAndServer(address string) {
	// bind listen address
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(fmt.Sprintf("listen err: %v", err))
		return
	}


	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal(fmt.Sprintf("close err: %v", err))
		}
	}(listener)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(fmt.Sprintf("accept err: %v", err))
		}

		go Handle(conn)
	}

}

func Handle(conn net.Conn) {
	render := bufio.NewReader(conn)
	for {
		msg, err := render.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("connection close")
			} else {
				log.Println(err)
			}
			return
		}

		b := [] byte(msg)
		conn.Write(b)
	}
}

func main() {
	ListenAndServer(":9999")
}
