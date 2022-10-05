package main

/*
	gRPC 的调用方式:
	- Unary RPC: 一元RPC
		客户端发起一次普通的RPC请求, 服务端响应一次请求, 是最基础的调用方式
	- Server-side streaming RPC: 服务端流式RPC
		单向流, 指Server为Stream, Client 为普通的一元RPC请求
	- Client-side streaming RPC: 客户端流式RPC
		单向流, 客户端通过流式发起多次RPC请求给服务端, 而服务端仅发起一次
		响应给客户端
	- Bidirectional streaming RPC: 双向流式RPC
		双向流, 由客户端以流式的方式发起请求, 服务端同样以流式的方式响应请求
	TODO: 熟悉以上调用方式的应用场景,  使用 wireshark 抓包分析客户端和服务端
		的交互过程
*/

import (
	"context"
	"flag"
	"io"
	"log"
	"net"

	pb "grpc-demo/proto"

	"google.golang.org/grpc"
)

type GreeterServer struct{}

// 一元RPC
func (s *GreeterServer) SayHello(ctx context.Context,
	r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello, world"}, nil
}

// 服务端流式RPC
func (s *GreeterServer) SayList(r *pb.HelloRequest,
	stream pb.Greeter_SayListServer) error {
	if r.Name != "client1" {
		return nil
	}
	for n := 0; n <= 6; n++ {
		// TODO: Send 方法源码
		_ = stream.Send(&pb.HelloReply{Message: "hello.list"})
	}
	return nil
}

// 客户端流式RPC
func (s *GreeterServer) SayRecord(stream pb.Greeter_SayRecordServer) error {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			message := &pb.HelloReply{Message: "say.record"}
			return stream.SendAndClose(message)
		}
		if err != nil {
			return err
		}

		log.Printf("resp: %v\n", resp)
	}
	// return nil
}

// 双向流式RPC
func (s *GreeterServer) SayRoute(stream pb.Greeter_SayRouteServer) error {
	n := 0
	for {
		_ = stream.Send(&pb.HelloReply{Message: "say.route"})

		resp, err := stream.Recv()
		if err == io.EOF {
			return err
		}
		if err != nil {
			return err
		}

		n++
		log.Printf("resp: %v\n", resp)
	}
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
