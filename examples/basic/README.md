# zdpgo_grpc

## 常用命令
生成proto
```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/helloworld.proto
protoc --proto_path=proto --go_out=proto proto/*.proto
```
