package main

import (
	"context"
	"log"
	pb "tag-service/proto"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	clientConn, _ := GetClientConn(ctx, "localhost:8004", nil)
	defer clientConn.Close()

	tagServiceClient := pb.NewTagServiceClient(clientConn)
	resp, _ := tagServiceClient.GetTagList(ctx,
		&pb.GetTagListRequest{Name: "kali"}) // TODO:  grpc.CallOption

	log.Printf("resp: %v\n", resp)
}

func GetClientConn(ctx context.Context, target string,
	opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure()) // TODO: WithInsecure
	return grpc.DialContext(ctx, target, opts...)
}
