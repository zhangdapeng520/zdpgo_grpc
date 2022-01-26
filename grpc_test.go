package zdpgo_grpc

import (
	"fmt"
	"testing"
)

func prepareGrpc() *Grpc {
	return New(GrpcConfig{
		Debug: true,
	})
}

func TestGrpc_New(t *testing.T) {
	g := prepareGrpc()
	fmt.Println(g)
}
