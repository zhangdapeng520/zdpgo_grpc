package g

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_password"
	"github.com/zhangdapeng520/zdpgo_yaml"
	"server/config"
)

// 导入此包会自动执行初始化方法
func init() {
	initConfig() // 初始化配置
	initAes()    // 初始化AES加密方法
	initEcc()
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

	// 默认加密算法：AES
	if Config.PasswordServer.EncryptAlgorithm == "" {
		Config.PasswordServer.EncryptAlgorithm = "aes"
	}
	// 默认AES加密的key
	if Config.PasswordServer.AesKey == "" {
		Config.PasswordServer.AesKey = "_ZhangDapeng520%"
	}
	// 默认AES加密的block块大小
	if Config.PasswordServer.AesBlockSize == 0 {
		Config.PasswordServer.AesBlockSize = 16
	}
}

// 初始化AES加密算法
func initAes() {
	Aes = zdpgo_password.NewAes(zdpgo_password.AesConfig{
		Key:       Config.PasswordServer.AesKey,
		BlockSize: Config.PasswordServer.AesBlockSize,
	})
}

// 初始化Ecc加密算法
func initEcc() {
	Ecc = zdpgo_password.NewEcc()
}
