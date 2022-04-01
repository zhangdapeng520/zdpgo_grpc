package config

type PasswordServerConfig struct {
	// 加密算法
	EncryptAlgorithm string `json:"encrypt_algorithm" yaml:"encrypt_algorithm"`
	AesKey           string `json:"aes_key" yaml:"aes_key"`
	AesBlockSize     int    `json:"aes_block_size" yaml:"aes_block_size"`
}

type Config struct {
	// 密码服务配置
	PasswordServer PasswordServerConfig `json:"password_server" yaml:"password_server"`
}
