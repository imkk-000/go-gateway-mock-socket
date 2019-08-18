package gateway

import "net"

type Gateway struct {
	ExternalServer net.Conn
}
