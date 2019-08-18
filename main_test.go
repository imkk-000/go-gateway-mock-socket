package main_test

import (
	. "mocking"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGatewayReceiveSuccess(t *testing.T) {
	message := "Receive From Mocking"
	gw := Gateway{
		ExternalServer: &MockSuccessExternalServer{},
	}

	data, n, _ := gw.Receive()

	assert.NotEmpty(t, data)
	assert.Equal(t, data, []byte(message))
	assert.Equal(t, len(message), n)
}

func TestGatewaySendSuccess(t *testing.T) {
	message := "Send To Mocking"
	gw := Gateway{
		ExternalServer: &MockSuccessExternalServer{},
	}

	n, _ := gw.Send([]byte(message))

	assert.Equal(t, len(message), n)
}
