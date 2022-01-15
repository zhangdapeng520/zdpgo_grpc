package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "zdpgo_grpc/examples/proto"
)

func main() {
	// 创建请求
	req := pb.HelloRequest{
		Name: "张大鹏",
	}

	// 解析请求，类似于json
	resp, err := proto.Marshal(&req)
	fmt.Println(resp, err)
}
