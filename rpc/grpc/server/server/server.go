package grpc_server

import (
	"context"
	"fmt"
	"github.com/aimfei/protocol/base_common"
	"github.com/aimfei/protocol/user_service"
	"google.golang.org/protobuf/types/known/anypb"
)

//Service 服务对象
type Service struct{}

func (s *Service) GetUserById(ctx context.Context, req *user_service.UserRequest) (*base_common.BaseResult, error) {
	id := req.GetUserId()
	fmt.Println(id)
	userResponse := &user_service.UserResponse{
		Username: "张三",
	}
	any, _ := anypb.New(userResponse)
	baseResult := &base_common.BaseResult{
		Code:    "00000",
		Message: "success",
		Data:    any,
	}
	return baseResult, nil
}
