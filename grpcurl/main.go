
func main() {
	s := grpc.NewServer()
	pb.RegisterGrpcServiceServer(s, &server{})
}