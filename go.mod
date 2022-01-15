module zdpgo_grpc

go 1.16

require (
	github.com/golang/protobuf v1.5.2
	github.com/satori/go.uuid v1.2.0
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
	zdpgo_consul v0.1.0
	zdpgo_gin v0.0.0
	zdpgo_nacos v0.1.0
	zdpgo_random v0.0.0
	zdpgo_zap v0.1.0
)

replace (
	zdpgo_config => ../zdpgo_config
	zdpgo_consul => ../zdpgo_consul
	zdpgo_gin => ../zdpgo_gin
	zdpgo_jwt => ../zdpgo_jwt
	zdpgo_nacos => ../zdpgo_nacos
	zdpgo_random => ../zdpgo_random
	zdpgo_zap => ../zdpgo_zap
)
