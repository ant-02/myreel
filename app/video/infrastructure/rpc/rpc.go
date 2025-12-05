package rpc

import (
	"context"
	"myreel/app/video/domain/repository"
	"myreel/kitex_gen/user"
	"myreel/kitex_gen/user/userservice"
	"myreel/pkg/errno"
)

type videoRpcImpl struct {
	user userservice.Client
}

func NewVideoRpcImpl(u userservice.Client) repository.RpcPort {
	return &videoRpcImpl{
		user: u,
	}
}

func (rpc *videoRpcImpl) GetUserIdByUsername(ctx context.Context, username string) (int64, error) {
	resp, err := rpc.user.GetUseridByUsername(ctx, &user.GetUserIdByUsernameRequest{
		Username: username,
	})
	if err != nil {
		return 0, errno.Errorf(errno.InternalRPCErrorCode, "rpc: failed to get user id by username: %v", err)
	}

	if resp.Base.Code != errno.SuccessCode {
		return 0, errno.NewErrNo(errno.InternalRPCErrorCode, "rpc: failed to get user id by username")
	}

	return resp.UserId, nil
}
