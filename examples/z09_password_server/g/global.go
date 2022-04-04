package g

import (
	"github.com/zhangdapeng520/zdpgo_password"
	"server/config"
)

var (
	Config   *config.Config           // 全局配置对象
	Password *zdpgo_password.Password // AES加密方法
)
