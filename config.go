package zdpgo_grpc

type Config struct {
	Debug       bool   `yaml:"debug" json:"debug"`                 // 是否为debug模式
	LogFilePath string `yaml:"log_file_path" json:"log_file_path"` // 日志路径
	ListenType  string `yaml:"listen_type" json:"listen"`          // 监听类型：默认tcp，还支持tcp4/tcp6/unix/unixpacket
	Server      Server `yaml:"server" json:"server"`               // 服务器配置
	Client      Server `yaml:"client" json:"client"`               // 客户端配置
}

type Server struct {
	Host string `yaml:"host" json:"host"` // 主机地址
	Port int    `yaml:"port" json:"port"` // 端口号
}
