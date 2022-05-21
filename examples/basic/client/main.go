package main

import (
	"basic/g"
	"context"
	"fmt"

	"basic/proto"
)

func main() {
	// 设置权限校验参数
	g.Grpc.Credential.Metadata = map[string]string{
		"username": "zhangdapeng111",
		"password": "zhangdapeng222",
	}

	// 连接grpc
	conn, err := g.Grpc.GetClientConn()
	if err != nil {
		panic(err)
	}
	defer g.Grpc.CloseClient(conn)

	// 创建客户端
	c := proto.NewGreeterClient(conn)

	// 发送请求，获取响应
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "张大鹏"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
