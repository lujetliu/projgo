package main

import (
	"context"
	"flag"
	pb "grpc-demo/proto"
	"io"
	"log"

	"google.golang.org/grpc"
)

var port string

func main() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	_ = SayHello(client)
}

// 一元RPC
func SayHello(client pb.GreeterClient) error {
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: "eddycjy"})
	log.Printf("client.SayHello resp: %s", resp.Message)
	return nil
}

// 服务端流式RPC
func SayList(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayList(context.Background(), r)
	for {
		// TODO: Recv 源码
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("io.EOF")
			break
		}
		if err != nil {
			log.Println(err)
			return err
		}

		log.Printf("resp: %v", resp)

	}

	return nil
}

// 客户端流式RPC
func SayRecord(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayRecord(context.Background())
	for n := 0; n < 6; n++ {
		_ = stream.Send(r)
	}
	resp, _ := stream.CloseAndRecv()

	log.Printf("resp err: %v", resp)
	return nil
}

// 双向流式RPC
func SayRoute(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayRoute(context.Background())
	for n := 0; n <= 6; n++ {
		_ = stream.Send(r)
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: %v\n", resp)
	}

	_ = stream.CloseSend()
	return nil
}
