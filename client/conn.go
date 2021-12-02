package client

import (
	"fmt"

	"google.golang.org/grpc"
)

// 获取GRPC的客户端连接
func GetConnect(ip string, port int) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure())
	return conn, err
}
