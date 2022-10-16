package main

import (
	"context"
	"log"
	pb "tag-service/proto"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	clientConn, _ := GetClientConn(ctx,
		"localhost:8001",
		[]grpc.DialOption{grpc.WithBlock()})
	defer clientConn.Close()

	tagServiceClient := pb.NewTagServiceClient(clientConn)
	resp, err := tagServiceClient.GetTagList(ctx,
		&pb.GetTagListRequest{Name: "kali"}) // TODO:  grpc.CallOption

	if err != nil {
		panic(err)
	}

	log.Printf("resp: %v\n", resp)
}

func GetClientConn(ctx context.Context, target string,
	opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	// 要请求的服务端是非加密模式的, 因此需调用 grpc.WithInsecure 方法禁用此
	// ClientConn 的传输安全性验证
	return grpc.DialContext(ctx, target, opts...)
}
