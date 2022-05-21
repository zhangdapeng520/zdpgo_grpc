package g

import (
	"github.com/zhangdapeng520/zdpgo_grpc"
)

/*
@Time : 2022/5/21 15:50
@Author : 张大鹏
@File : initialize.go
@Software: Goland2021.3.1
@Description: 初始化全局变量
*/

func init() {
	InitGrpc()
}

func InitGrpc() {
	Grpc = zdpgo_grpc.NewWithConfig(&zdpgo_grpc.Config{
		Debug:       true,
		LogFilePath: "",
		Server: zdpgo_grpc.Server{
			Host: "127.0.0.1",
			Port: 33333,
		},
		Client: zdpgo_grpc.Server{
			Host: "127.0.0.1",
			Port: 33333,
		},
	})
}
