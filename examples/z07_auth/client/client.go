package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"zdpgo_grpc/examples/proto"
)

type customCredential struct{}

// 设置metadata
func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	// 返回map结构，map中存放自定义的metadata信息
	return map[string]string{
		"appid":  "zhangdapeng",
		"appkey": "zhangdapeng520", // 可以修改值，体验验证失败
	}, nil
}

// 是否使用安全协议
func (c customCredential) RequireTransportSecurity() bool {
	return false
}

func main() {
	// 使用自定义的协议
	grpc.WithPerRPCCredentials(customCredential{})

	// 请求要携带的metadata信息
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithPerRPCCredentials(customCredential{}))

	// 建立连接
	conn, err := grpc.Dial("127.0.0.1:50051", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 创建客户端
	c := proto.NewGreeterClient(conn)

	// 发送请求
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "zhangdapeng"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
