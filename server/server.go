package server

import (
	"context"
	"net"

	"github.com/wdfky/grpcdemo/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GreeterServiceImpl struct {
	pb.UnimplementedGreeterServer
}

func (g *GreeterServiceImpl) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "hello" + req.GetName(),
	}, nil
}
func (*GreeterServiceImpl) mustEmbedUnimplementedGreeterServer() {}

func Init() {
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, new(GreeterServiceImpl))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	reflection.Register(s)
	s.Serve(listener)
}
