- 构造
```
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --grpc-gateway_out=./gateway --go_out=plugins=grpc:./gateway hello.proto

```

- 启动
```
<!-- 启动gateway -->
go run main.go
<!-- 启动 grpc server -->
go run server/server.go
<!-- 浏览器请求 -->
http://localhost:8080/get/gopher
<!-- 终端 -->
curl localhost:8080/post -X POST --data '{"value":"grpc"}'
curl localhost:8080/get/go

```