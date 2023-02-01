package server

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"test_grpc/pb"
)

var (
	addr = "127.0.0.1:1998"
)

type Server struct{}

func StartServe() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("[StartServe] Listen err", err.Error())
		return
	}

	s := grpc.NewServer()
	pb.RegisterGreetServer(s, &Server{})

	reflection.Register(s)

	fmt.Println("rpc服务端已启动～")
	if err = s.Serve(lis); err != nil {
		fmt.Println("[StartServe] Serve err", err.Error())
	}
}
