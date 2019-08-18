package gateway_test

import (
	"bufio"
	. "mocking/gateway"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGatewayWithMock(t *testing.T) {
	expectedResult := []byte("Response From Gateway Mocking")
	gateway := Gateway{
		ExternalServer: &MockExternalServer{},
	}

	reader := bufio.NewReader(gateway.ExternalServer)
	actualResult := make([]byte, 2048)
	n, _ := reader.Read(actualResult)

	assert.Equal(t, expectedResult, actualResult[:n])
}
