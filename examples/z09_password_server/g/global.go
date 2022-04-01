package g

import (
	"github.com/zhangdapeng520/zdpgo_password"
	"server/config"
)

var (
	Config *config.Config      // 全局配置对象
	Aes    *zdpgo_password.Aes // AES加密方法
	Ecc    *zdpgo_password.Ecc // ECC加密方法
)
