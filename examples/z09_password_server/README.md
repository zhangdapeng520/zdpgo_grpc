# 密码加密解密的服务

## 生成proto代码
```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative password.proto
```