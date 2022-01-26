package main

import (
	"context"
	"fmt"
	"github.com/aimfei/protocol/user_service"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:1090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("start error %+v", err)
	}
	defer conn.Close()
	client := user_service.NewUserServiceClient(conn)
	userRequest := &user_service.UserRequest{
		UserId: 1,
	}
	baseResult, err := client.GetUserById(context.Background(), userRequest)
	if err != nil {
		log.Fatalf("rpc baseResult error %+v", err)
	}
	var userInfo user_service.UserResponse
	err = anypb.UnmarshalTo(baseResult.GetData(), &userInfo, proto.UnmarshalOptions{})
	fmt.Printf("%+v\n", baseResult)
	fmt.Printf("%s", userInfo.GetUsername())
}
