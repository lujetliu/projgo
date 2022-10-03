package main

import (
	"context"
	"flag"
	"net"

	pb "grpc-demo/proto"

	"google.golang.org/grpc"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello, world"}, nil
}

var port string

func main() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()

	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, _ := net.Listen("tcp", ":"+port)
	server.Serve(lis)
}
