package zdpgo_grpc

import "context"

/*
@Time : 2022/5/21 17:02
@Author : 张大鹏
@File : credential.go
@Software: Goland2021.3.1
@Description: credential 安全协议相关
*/

type Credential struct {
	Metadata map[string]string // 校验数据
}

// GetRequestMetadata 设置metadata
func (c Credential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	// 返回map结构，map中存放自定义的metadata信息
	if c.Metadata == nil {
		c.Metadata = map[string]string{
			"username": "zhangdapeng",
			"password": "zhangdapeng", // 可以修改值，体验验证失败
		}
	}
	return c.Metadata, nil
}

// RequireTransportSecurity 是否使用安全协议
func (c Credential) RequireTransportSecurity() bool {
	return false
}
