package main

import (
	"fmt"
	"net"

	"github.com/mytaowu/grpc-server/service"
	pb "github.com/mytaowu/proto"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello world")
	list, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	server := grpc.NewServer()
	pb.RegisterHelloGRPCServer(server, &service.HelloService{})

	fmt.Println("grpc 服务启动")
	server.Serve(list)
}
