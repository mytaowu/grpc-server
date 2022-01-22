package service

import (
	"context"
	"fmt"

	pb "github.com/mytaowu/proto"
)

type HelloService struct {
	pb.UnimplementedHelloGRPCServer
}

func (h *HelloService) SayHi(ctx context.Context, req *pb.Req) (*pb.Res, error) {
	fmt.Println("hello, world: msg: ", req.Message)

	return &pb.Res{
		Message: "say hi" + req.Message,
	}, nil
}
