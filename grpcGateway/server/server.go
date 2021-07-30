package main

import (
	"context"
	"net"

	"studyGo/grpcGateway/gateway"

	"google.golang.org/grpc"
)

type RestServiceImpl struct{}

func (r *RestServiceImpl) Get(ctx context.Context, message *gateway.StringMessage) (*gateway.StringMessage, error) {
	return &gateway.StringMessage{Value: "Get hi:" + message.Value + "#"}, nil
}

func (r *RestServiceImpl) Post(ctx context.Context, message *gateway.StringMessage) (*gateway.StringMessage, error) {
	return &gateway.StringMessage{Value: "Post hi:" + message.Value + "@"}, nil
}
func main() {
	grpcServer := grpc.NewServer()
	gateway.RegisterRestServiceServer(grpcServer, new(RestServiceImpl))
	lis, _ := net.Listen("tcp", ":5000")
	grpcServer.Serve(lis)
}
