module zgo_grpc

go 1.16

require (
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
	zgo_api v0.1.0
	zgo_consul v0.1.0
	zgo_log v0.1.0
	zgo_nacos v0.1.0
	zgo_random v0.1.0
)
replace (
	zgo_api => ../zgo_api
	zgo_jwt => ../zgo_jwt
	zgo_config => ../zgo_config
	zgo_consul => ../zgo_consul
	zgo_log => ../zgo_log
	zgo_nacos => ../zgo_nacos
	zgo_random => ../zgo_random
)
