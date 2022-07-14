package zdpgo_grpc

import (
	"errors"

	"google.golang.org/grpc"
)

/*
@Time : 2022/5/21 15:47
@Author : 张大鹏
@File : close.go
@Software: Goland2021.3.1
@Description: close类型的关闭方法
*/

func (g *Grpc) CloseClient(conn *grpc.ClientConn) error {
	// 不需要关闭
	if conn == nil {
		return errors.New("client connection is nil")
	}

	// 关闭连接
	err := conn.Close()
	if err != nil {
		return err
	}

	return nil
}
