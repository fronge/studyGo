package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "studyGo/grpcServer/grpcT"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port = ":50051"
)

type server struct { //服务的结构类型
	*pb.UnimplementedGrpcServiceServer
}

func (s *server) Fun(ctx context.Context, in *pb.RequestData) (*pb.ResponseData, error) {
	fmt.Println("-------%d,===%s", in.R, in.RepT)
	return &pb.ResponseData{ResT: "aaa", Code: 200}, nil
}

func (s *server) A(ctx context.Context, in *pb.RequestData) (*pb.ResponseData, error) {
	fmt.Println("-------%d,===%s", in.R, in.RepT)
	return &pb.ResponseData{ResT: "aaa", Code: 200}, nil
}

func main() {
	creds, err := credentials.NewServerTLSFromFile("server.pem", "server.key")
	s := grpc.NewServer(grpc.Creds(creds))

	lis, err := net.Listen("tcp", port) //开启监听
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// s := grpc.NewServer()                      //新建一个grpc服务
	pb.RegisterGrpcServiceServer(s, &server{}) //这个服务和上述的服务结构联系起来，这样你新建的这个服务里面就有那些类型的方法
	if err := s.Serve(lis); err != nil {       //这个服务和你的监听联系起来，这样外界才能访问到啊
		log.Fatalf("failed to serve: %v", err)
	}
}
