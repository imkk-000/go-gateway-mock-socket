package main

import (
	"log"
	"net"
)

type Gateway struct {
	ExternalServer net.Conn
}

func (gw Gateway) Send(data []byte) (n int, err error) {
	n, err = gw.ExternalServer.Write(data)
	return
}

func (gw Gateway) Receive() (data []byte, n int, err error) {
	data = make([]byte, 2048)
	n, err = gw.ExternalServer.Read(data)
	data = data[:n]
	return
}

func newClient() net.Conn {
	conn, err := net.Dial("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func main() {
	gw := Gateway{
		ExternalServer: newClient(),
	}

	// send
	n, err := gw.Send([]byte("Hello Server"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Send data size", n)

	// receive
	b, n, err := gw.Receive()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(n, b)
}
