package zgo_grpc

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"zgo_consul"
	"zgo_log"
	"zgo_nacos"
	"zgo_random/http"
	"zgo_random/uuid"
)

// 启动grpc service服务
// 参数1：要启动的grpc服务
func StartGrpcService(server *grpc.Server) {
	var err error

	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 0, "端口号")

	// 动态获取命令行参数
	flag.Parse()

	// 动态获取端口号
	if *Port == 0 {
		*Port, err = http.FreePort()
		zgo_nacos.ServiceConfig.Port = *Port
		if err != nil {
			zgo_log.L().Panic("动态获取端口号失败", zgo_log.String("err", err.Error()))
		}
	}
	zgo_log.L().Info("初始化信息", zgo_log.String("ip", *IP), zgo_log.Int("port", *Port))
	zgo_log.L().Info("nacos配置信息", zgo_log.Any("nacos", zgo_nacos.ServiceConfig))

	// 创建并监听grpc服务
	address := fmt.Sprintf("%s:%d", *IP, *Port)
	zgo_log.L().Info("监听tcp端口号", zgo_log.String("address", address))
	lis, err := net.Listen("tcp", address)
	if err != nil {
		zgo_log.L().Panic("监听grpc服务失败", zgo_log.String("err", err.Error()))
	}

	//注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	//服务注册
	consulClient := zgo_consul.GetConsulClient(
		zgo_nacos.ServiceConfig.Consul.Host,
		zgo_nacos.ServiceConfig.Consul.Port,
	)
	serviceID := fmt.Sprintf("%s", uuid.NewV4())
	zgo_consul.RegisterGrpc(
		consulClient,
		zgo_nacos.ServiceConfig.Host,
		zgo_nacos.ServiceConfig.Port,
		zgo_nacos.ServiceConfig.Name,
		serviceID,
		zgo_nacos.ServiceConfig.Tags,
	)

	zgo_log.L().Info("服务注册到consul成功")
	zgo_log.L().Info("启动服务")
	go func() {
		err = server.Serve(lis)
		if err != nil {
			zgo_log.L().Fatal("启动服务失败", zgo_log.String("err", err.Error()))
		}
	}()

	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zgo_log.L().Info("停止服务")
	zgo_consul.DeRegister(consulClient, serviceID)
	zgo_log.L().Info("注销微服务成功")
}
