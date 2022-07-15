package main

import (
	"context"
	"fmt"

	"github.com/zhangdapeng520/zdpgo_grpc/examples/token/proto"

	"google.golang.org/grpc"
)

// 自定义权限校验对象，必须实现以下两个方法
type customCredential struct{}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	// 传递appid和appkey
	return map[string]string{
		"appid":  "101010",
		"appkey": "i am key",
	}, nil
}

func (c customCredential) RequireTransportSecurity() bool {
	return false
}

func main() {
	// 创建自定义权限校验器
	grpc.WithPerRPCCredentials(customCredential{})

	// 使用权限校验器
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithPerRPCCredentials(customCredential{}))

	// 创建连接
	conn, err := grpc.Dial("127.0.0.1:50051", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 创建客户端
	c := proto.NewGreeterClient(conn)

	// 发送请求
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "bobby"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
