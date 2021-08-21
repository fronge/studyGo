package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "studyGo/grpcServer/grpcT"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	// "google.golang.org/grpc/internal/metadata"
	"google.golang.org/grpc/metadata"
)

const (
	port = ":50051"
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

type server struct { //服务的结构类型
	*pb.UnimplementedGrpcServiceServer
	auth *Authentication
}

func (s *server) Fun(ctx context.Context, in *pb.RequestData) (*pb.ResponseData, error) {
	if err := s.auth.Auth(ctx); err != nil {
		return nil, err
	}
	return &pb.ResponseData{ResT: "aaa", Code: 200}, nil
}

func (s *server) A(ctx context.Context, in *pb.RequestData) (*pb.ResponseData, error) {
	fmt.Println("-------%d,===%s", in.R, in.RepT)
	return &pb.ResponseData{ResT: "aaa", Code: 200}, nil
}

func (a *Authentication) Auth(ctx context.Context) error {
	// 提取出元信息
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Println(md)
	if !ok {
		return fmt.Errorf("missing credentials")
	}

	var appid string
	var appkey string
	if val, ok := md["user"]; ok {
		appid = val[0]
	}

	if val, ok := md["password"]; ok {
		appkey = val[0]
	}
	fmt.Println(appid)
	fmt.Println(appkey)
	if appid != a.User || appkey != a.Password {
		return fmt.Errorf("codes Unathenticated invalidtoken")
	}

	return nil
}

func filter(ctx context.Context,
	req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	log.Println("fileter:", info)
	return handler(ctx, req)
}

func main() {
	creds, err := credentials.NewServerTLSFromFile("server.pem", "server.key")
	s := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(filter))

	lis, err := net.Listen("tcp", port) //开启监听
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	auth := Authentication{
		User:     "gopher",
		Password: "password",
	}
	// s := grpc.NewServer()                      //新建一个grpc服务
	pb.RegisterGrpcServiceServer(s, &server{
		auth: &auth,
	}) //这个服务和上述的服务结构联系起来，这样你新建的这个服务里面就有那些类型的方法
	if err := s.Serve(lis); err != nil { //这个服务和你的监听联系起来，这样外界才能访问到啊
		log.Fatalf("failed to serve: %v", err)
	}
}

// select uuid as baidu_std from position_details_baidu_std where company_name = "上海玛帕贸易有限公司";
// select uuid as chinahr from position_details_chinahr where company_name = "上海玛帕贸易有限公司";
// select uuid as cjol from position_details_cjol where company_name = "上海玛帕贸易有限公司";
// select uuid as dajie from position_details_dajie where company_name = "上海玛帕贸易有限公司";
// select uuid as haitou from position_details_haitou where company_name = "上海玛帕贸易有限公司";
// select uuid as huangye88 from position_details_huangye88 where company_name = "上海玛帕贸易有限公司";
// select uuid as itjuzi from position_details_itjuzi where company_name = "上海玛帕贸易有限公司";
// select uuid as jyjob from position_details_jyjob where company_name = "上海玛帕贸易有限公司";
// select uuid as kanzhun from position_details_kanzhun where company_name = "上海玛帕贸易有限公司";
// select uuid as lagou from position_details_lagou where company_name = "上海玛帕贸易有限公司";
// select uuid as liepin from position_details_liepin where company_name = "上海玛帕贸易有限公司";
// select uuid as meiye from position_details_meiye where company_name = "上海玛帕贸易有限公司";
// select uuid as qimingpian from position_details_qimingpian where company_name = "上海玛帕贸易有限公司";
// select uuid as qiyeminglu from position_details_qiyeminglu where company_name = "上海玛帕贸易有限公司";
// select uuid as shunqi from position_details_shunqi where company_name = "上海玛帕贸易有限公司";
// select uuid as tuilie from position_details_tuilie where company_name = "上海玛帕贸易有限公司";
// select uuid as wuba from position_details_wuba where company_name = "上海玛帕贸易有限公司";
// select uuid as wuyou from position_details_wuyou where company_name = "上海玛帕贸易有限公司";
// select uuid as yingcai from position_details_yingcai where company_name = "上海玛帕贸易有限公司";
// select uuid as zhaopingou from position_details_zhaopingou where company_name = "上海玛帕贸易有限公司";
// select uuid as zhilian from position_details_zhilian where company_name = "上海玛帕贸易有限公司";
// select uuid as zhipin from position_details_zhipin where company_name = "上海玛帕贸易有限公司";
// select uuid as zhiq from position_details_zhiq where company_name = "上海玛帕贸易有限公司";
// select uuid as zyquan7360 from position_details_zyquan7360 where company_name = "上海玛帕贸易有限公司";