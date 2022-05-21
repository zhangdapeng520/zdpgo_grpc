package zdpgo_grpc

import (
	"github.com/zhangdapeng520/zdpgo_log"
)

// Grpc 核心对象
type Grpc struct {
	Log        *zdpgo_log.Log // 日志对象
	Config     *Config        // 配置对象
	Credential *Credential    // 自定义校验对象
}

func New() *Grpc {
	return NewWithConfig(&Config{})
}

func NewWithConfig(config *Config) *Grpc {
	g := &Grpc{}

	// 日志
	if config.LogFilePath == "" {
		config.LogFilePath = "logs/zdpgo/zdpgo_grpc.log"
	}
	g.Log = zdpgo_log.NewWithDebug(config.Debug, config.LogFilePath)

	// 配置
	if config.Server.Host == "" {
		config.Server.Host = "0.0.0.0"
	}
	if config.Server.Port == 0 {
		config.Server.Port = 33333
	}
	if config.Client.Host == "" {
		config.Client.Host = "localhost"
	}
	if config.Client.Port == 0 {
		config.Client.Port = 33333
	}
	g.Config = config

	// 自定义校验对象
	g.Credential = &Credential{}

	// 返回
	return g
}
