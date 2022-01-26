package main

import (
	"github.com/aimfei/protocol/user_service"
	grpc_server "golang-awesome/rpc/grpc/server/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:1090")
	if err != nil {
		log.Fatalf("%+v", err)
	}
	grpcServer := grpc.NewServer()
	userService := grpc_server.Service{}

	user_service.RegisterUserServiceServer(grpcServer, &userService)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Fatalln("grpc_server start success")
}
