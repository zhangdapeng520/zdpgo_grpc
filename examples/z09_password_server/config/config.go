package config

import "github.com/zhangdapeng520/zdpgo_password"

type Config struct {
	// 密码配置
	PasswordConfig zdpgo_password.PasswordConfig `json:"password_config" yaml:"password_config"`
}
