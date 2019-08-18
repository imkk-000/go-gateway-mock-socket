package gateway_test

import (
	"net"
	"time"
)

type MockExternalServer struct{}

func (MockExternalServer) Read(b []byte) (n int, err error) {
	n = copy(b, []byte("Response From Gateway Mocking"))
	return n, nil
}
func (MockExternalServer) Write(b []byte) (n int, err error)  { return 0, nil }
func (MockExternalServer) Close() error                       { return nil }
func (MockExternalServer) LocalAddr() net.Addr                { return nil }
func (MockExternalServer) RemoteAddr() net.Addr               { return nil }
func (MockExternalServer) SetDeadline(t time.Time) error      { return nil }
func (MockExternalServer) SetReadDeadline(t time.Time) error  { return nil }
func (MockExternalServer) SetWriteDeadline(t time.Time) error { return nil }
