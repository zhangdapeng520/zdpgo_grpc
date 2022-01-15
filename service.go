package zdpgo_grpc

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"zdpgo_consul"
	"zdpgo_nacos"
	zdpgo_random "zdpgo_random"
	"zdpgo_zap"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// 启动grpc service服务
// 参数1：要启动的grpc服务
func StartGrpcService(server *grpc.Server) {
	var err error
	var random = zdpgo_random.New()
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 0, "端口号")

	// 动态获取命令行参数
	flag.Parse()

	// 动态获取端口号
	if *Port == 0 {
		*Port = random.RandomHttpPort()
		zdpgo_nacos.ServiceConfig.Port = *Port
	}
	zdpgo_zap.L().Info("初始化信息", zdpgo_zap.String("ip", *IP), zdpgo_zap.Int("port", *Port))
	zdpgo_zap.L().Info("nacos配置信息", zdpgo_zap.Any("nacos", zdpgo_nacos.ServiceConfig))

	// 创建并监听grpc服务
	address := fmt.Sprintf("%s:%d", *IP, *Port)
	zdpgo_zap.L().Info("监听tcp端口号", zdpgo_zap.String("address", address))
	lis, err := net.Listen("tcp", address)
	if err != nil {
		zdpgo_zap.L().Panic("监听grpc服务失败", zdpgo_zap.String("err", err.Error()))
	}

	//注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	//服务注册
	consulClient := zdpgo_consul.GetConsulClient(
		zdpgo_nacos.ServiceConfig.Consul.Host,
		zdpgo_nacos.ServiceConfig.Consul.Port,
	)
	serviceID := fmt.Sprintf("%s", uuid.NewV4())
	fmt.Println("注册GRPC健康检查：", zdpgo_nacos.ServiceConfig.Host, zdpgo_nacos.ServiceConfig.Port)
	zdpgo_consul.RegisterGrpc(
		consulClient,
		zdpgo_nacos.ServiceConfig.Host,
		zdpgo_nacos.ServiceConfig.Port,
		zdpgo_nacos.ServiceConfig.Name,
		serviceID,
		zdpgo_nacos.ServiceConfig.Tags,
	)

	zdpgo_zap.L().Info("服务注册到consul成功")
	zdpgo_zap.L().Info("启动服务")
	go func() {
		err = server.Serve(lis)
		if err != nil {
			zdpgo_zap.L().Fatal("启动服务失败", zdpgo_zap.String("err", err.Error()))
		}
	}()

	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zdpgo_zap.L().Info("停止服务")
	zdpgo_consul.DeRegister(consulClient, serviceID)
	zdpgo_zap.L().Info("注销微服务成功")
}
