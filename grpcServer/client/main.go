package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "studyGo/grpcServer/grpcT"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address = "localhost:50051"
)

type PerRPCCredentials interface {
	GetRequestMetadata(ctx context.Context, uri ...string) (
		map[string]string, error,
	)
	RequireTransportSecurity() bool
}

type Authentication struct {
	User     string
	Password string
}

func (a *Authentication) GetRequestMetadata(
	context.Context,
	...string) (
	map[string]string,
	error,
) {
	return map[string]string{"user": a.User, "password": a.Password}, nil
}

func (a *Authentication) RequireTransportSecurity() bool {
	return false
}

func main() {
	auth := Authentication{
		User:     "gopher",
		Password: "password",
	}
	// conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	creds, err := credentials.NewClientTLSFromFile("server.pem", "test")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatalf("connect to server failed: %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcServiceClient(conn) //返回一个client连接，通过这个连接就可以访问到对应的服务资源，就像一个对象
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second) //返回一个client，并设置超时时间
	defer cancel()
	r, err := c.Fun(ctx, &pb.RequestData{RepT: "aaa", R: int64(10)}) //访问对应的服务器上面的服务方法
	if err != nil {
		log.Fatalf("could not rpc: %v,%T", err, c)
	}
	fmt.Println(fmt.Sprintf("%v", r))
}
