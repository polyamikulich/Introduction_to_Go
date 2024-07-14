package main

import (
	"Go_Project/grpc/server"
	"Go_Project/proto"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {

	Server := server.New()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	proto.RegisterAccountsServer(s, Server)
	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
