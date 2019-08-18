package main_test

import (
	"net"
	"time"
)

type MockSuccessExternalServer struct{}

func (MockSuccessExternalServer) Read(b []byte) (n int, err error) {
	n = copy(b, []byte("Receive From Mocking"))
	return n, nil
}
func (MockSuccessExternalServer) Write(b []byte) (n int, err error) {
	return len(b), nil
}
func (MockSuccessExternalServer) Close() error                       { return nil }
func (MockSuccessExternalServer) LocalAddr() net.Addr                { return nil }
func (MockSuccessExternalServer) RemoteAddr() net.Addr               { return nil }
func (MockSuccessExternalServer) SetDeadline(t time.Time) error      { return nil }
func (MockSuccessExternalServer) SetReadDeadline(t time.Time) error  { return nil }
func (MockSuccessExternalServer) SetWriteDeadline(t time.Time) error { return nil }
