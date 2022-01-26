package zdpgo_grpc

import (
	"github.com/zhangdapeng520/zdpgo_zap"
)

// Grpc 核心对象
type Grpc struct {
	log    *zdpgo_zap.Zap // 日志对象
	config *GrpcConfig    // 配置对象
}

// GrpcConfig 配置
type GrpcConfig struct {
	Debug       bool   // 是否为debug模式
	LogFilePath string // 日志路径
}

func New(config GrpcConfig) *Grpc {
	g := Grpc{}
	
	// 日志
	if config.LogFilePath == "" {
		config.LogFilePath = "zdpgo_grpc.log"
	}
	g.log = zdpgo_zap.New(zdpgo_zap.ZapConfig{
		Debug:        config.Debug,
		OpenGlobal:   true,
		OpenFileName: true,
		LogFilePath:  config.LogFilePath,
	})
	
	// 配置
	g.config = &config
	
	return &g
}
