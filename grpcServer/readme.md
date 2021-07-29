
- 安装依赖
```
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
brew install protobuf
go get -u google.golang.org/grpc
```
- 生成代码
```
protoc --go_out=plugins=grpc:. grpcT.proto
```

- 生成秘钥
  - 注意 common name 选项要添加，且记录下来，本例子用的是test
```
openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650
```

- 运行 server
```
 go run server/main.go
 ```

- 运行 client
```
 go run client/main.go
 ```