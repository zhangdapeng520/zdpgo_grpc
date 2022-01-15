package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "zdpgo_grpc/examples/proto"
)

type Hello struct {
	Name string `json:"name"`
}

func main() {
	// 创建请求
	req := pb.HelloRequest{
		Name: "张大鹏",
	}

	// 解析请求，类似于json
	resp, err := proto.Marshal(&req)
	fmt.Println(string(resp), err)

	// 与json比较
	req1 := Hello{Name: "张大鹏"}
	resp1, err := json.Marshal(req1)
	fmt.Println(string(resp1), err)

	// 查看长度
	fmt.Println(len(resp1), len(resp), float32(len(resp1))/float32(len(resp)))
}
