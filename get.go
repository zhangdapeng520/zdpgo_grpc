package zdpgo_grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

/*
@Time : 2022/5/21 15:46
@Author : 张大鹏
@File : get.go
@Software: Goland2021.3.1
@Description: get类型的获取方法
*/

// GetClientConn 获取GRPC客户端连接
func (g *Grpc) GetClientConn() (conn *grpc.ClientConn, err error) {
	// 使用自定义的协议
	grpc.WithPerRPCCredentials(*g.Credential)

	// 请求要携带的metadata信息
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithPerRPCCredentials(*g.Credential))

	// 建立连接
	address := fmt.Sprintf("%s:%d", g.Config.Client.Host, g.Config.Client.Port)
	conn, err = grpc.Dial(address, opts...)
	if err != nil {
		g.Log.Error("获取客户端连接失败", "error", err, "address", address)
	}
	g.Log.Debug("获取GRPC客户端成功", "address", address)

	// 创建具体的GRPC客户端
	return
}

// GetInterceptor 获取拦截器
func (g *Grpc) GetInterceptor(beforeFunc func(ctx context.Context) error, afterFunc func(ctx context.Context)) grpc.ServerOption {
	// 创建一个拦截器
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// 执行拦截之前的方法
		if beforeFunc != nil {
			err = beforeFunc(ctx)
			if err != nil {
				return resp, err
			}
		}

		// 执行请求
		res, err := handler(ctx, req)

		// 执行拦截之后的方法
		if afterFunc != nil {
			afterFunc(ctx)
		}

		// 返回响应
		return res, err
	}

	// 使用拦截器
	opt := grpc.UnaryInterceptor(interceptor)

	return opt
}

// GetAuthInterceptor 设置权限校验拦截器
func (g *Grpc) GetAuthInterceptor(username, password string) grpc.ServerOption {
	return g.GetInterceptor(func(ctx context.Context) error {
		// 从拦截器中获取metadata
		if metaData, ok := metadata.FromIncomingContext(ctx); ok {
			// 从拦截器中获取token认证信息
			var (
				metaUsername string
				metaPassword string
			)
			// 获取用户名
			if usernameList, ok := metaData["username"]; ok {
				metaUsername = usernameList[0]
			}

			// 获取密码
			if passwordList, ok := metaData["password"]; ok {
				metaPassword = passwordList[0]
			}

			// 校验token认证信息
			if username != metaUsername || password != metaPassword {
				return status.Error(codes.Unauthenticated, "TOKEN校验失败，用户名或密码错误")
			}

			// 校验成功
			g.Log.Debug("权限校验成功", "username", username)
			return nil
		}
		return status.Error(codes.Unauthenticated, "无token认证信息")

	}, nil)
}
