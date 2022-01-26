package zdpgo_grpc

// ServerConfig 服务配置
type ServerConfig struct {
	Host string // 主机地址
	Port uint16 // 端口号
}

// ClientConfig 客户端配置
type ClientConfig struct {
	Host string // 主机地址
	Port uint16 // 端口号
}
