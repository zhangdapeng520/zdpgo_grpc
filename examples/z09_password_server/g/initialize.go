package g

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_password"
	"github.com/zhangdapeng520/zdpgo_yaml"
	"server/config"
)

// 导入此包会自动执行初始化方法
func init() {
	initConfig()   // 初始化配置
	initPassword() // 初始化AES加密方法
}

// 初始化配置
func initConfig() {
	// 创建yaml对象
	y := zdpgo_yaml.New()

	// 实例化配置
	Config = &config.Config{}

	// 读取配置
	err := y.ReadDefaultConfig(Config)
	if err != nil {
		fmt.Println("读取配置文件失败：", err)
	}
}

// 初始化AES加密算法
func initPassword() {
	Password = zdpgo_password.New(Config.PasswordConfig)
}
