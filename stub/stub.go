package main

import (
	"io"
	"log"
	"net"
)

func main() {
	stubServer, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer stubServer.Close()

	for {
		client, err := stubServer.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			defer c.Close()
			io.Copy(c, c)
			c.Write([]byte("\n"))
		}(client)
	}
}
