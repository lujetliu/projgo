package main

/*
	gRPC 是基于HTTP/2协议的, 所以不能用常用的接口测试工具进行测试(postman),
	可以使用 grpcurl 工具调试 gRPC 接口
	go get github.com/fullstorydev/grpcurl
	使用该工具的前提是 gRPC Server 注册了反射服务(Line31), reflection 包是
	gRPC 官方提供的反射服务, 在启动文件中新增 reflection.Register 方法的调用;

	调试:
	列出服务
	grpcurl -plaintext localhost:8001 list
	TagService
	grpc.reflection.v1alpha.ServerReflection

	查看服务的方法
	grpcurl -plaintext localhost:8001 list TagService
	TagService.GetTagList

	调用RPC方法
	grpcurl -plaintext -d '{"name":"kali"}' localhost:8001  TagService.GetTagList
	{
	  "list": [
		{
		  "id": "2",
		  "name": "kali",
		  "state": 1
		}
	  ],
	  "pager": {
		"page": "1",
		"pageSize": "10",
		"totalRows": "1"
	  }
	}

	grpcurl 工具的选项:
	- plaintext: 忽略TLS认证
	- list: 指定所执行的命令, list 子命令可获取该服务的RPC方法列表信息
*/

import (
	"flag"
	"log"
	"net"
	pb "tag-service/proto"
	"tag-service/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port string

func main() {
	flag.StringVar(&port, "p", "8001", "启动端口号")
	flag.Parse()

	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s) // TODO: 注册反射服务的作用

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Printf("net.Listen err: %v\n", err)
		return
	}

	err = s.Serve(lis)
	if err != nil {
		log.Printf("server.Server err %v\n", err)
	}
}
