package main

import (
	"context"
	"flag"
	pb "grpc-demo/proto"
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

func SayHello(client pb.GreeterClient) error {
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: "eddycjy"})
	log.Printf("client.SayHello resp: %s", resp.Message)
	return nil
}
