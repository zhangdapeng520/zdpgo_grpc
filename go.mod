module zdpgo_grpc

go 1.16

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/golang/protobuf v1.5.2
	google.golang.org/grpc v1.44.0
	google.golang.org/protobuf v1.27.1
	github.com/zhangdapeng520/zdpgo_gin v0.1.0
	github.com/zhangdapeng520/zdpgo_zap v0.1.0
	github.com/zhangdapeng520/zdpgo_code v0.1.0
)

replace (
	github.com/zhangdapeng520/zdpgo_gin v0.1.0 => ../zdpgo_gin
	github.com/zhangdapeng520/zdpgo_zap v0.1.0 => ../zdpgo_zap
	github.com/zhangdapeng520/zdpgo_code v0.1.0 => ../zdpgo_code
)
