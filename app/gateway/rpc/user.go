package rpc

import (
	"context"
	"myreel/kitex_gen/user"
	"myreel/kitex_gen/user/userservice"
	"myreel/pkg/constants"
	"myreel/pkg/errno"
	"myreel/pkg/logger"
	"myreel/pkg/util"

	"github.com/cloudwego/kitex/client"
)

func InitUserClient() {
	c, err := userservice.NewClient(constants.UserServiceName, client.WithHostPorts("0.0.0.0:20002"))
	if err != nil {
		logger.Fatalf("api.rpc.user InitUserRPC failed, err is %v", err)
	}
	userClient = c
}

func RegisterRPC(ctx context.Context, req *user.RegisterRequest) error {
	resp, err := userClient.Register(ctx, req)
	if err != nil {
		logger.Errorf("RegisterRPC: RPC called failed: %v", err.Error())
		return errno.InternalServiceError.WithError(err)
	}
	if !util.IsSuccess(resp.Base) {
		return errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}

	return nil
}
