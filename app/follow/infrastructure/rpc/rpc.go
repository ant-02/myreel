package rpc

import (
	"context"
	"myreel/app/follow/domain/model"
	"myreel/app/follow/domain/repository"
	"myreel/kitex_gen/user"
	"myreel/kitex_gen/user/userservice"
	"myreel/pkg/errno"
)

type followRpcImpl struct {
	user userservice.Client
}

func NewFollowRpcImpl(u userservice.Client) repository.RpcPort {
	return &followRpcImpl{
		user: u,
	}
}

func (rpc *followRpcImpl) GetUsersByIdsRPC(ctx context.Context, ids []int64) ([]*model.UserProfile, error) {
	resp, err := rpc.user.GetUsersByIds(ctx, &user.GetUsersByIdsRequest{
		Ids: ids,
	})
	if err != nil {
		return nil, errno.Errorf(errno.InternalRPCErrorCode, "rpc: failed to get users by ids: %v", err)
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(errno.InternalRPCErrorCode, "rpc: failed to get users by ids")
	}

	l := len(resp.List)
	users := make([]*model.UserProfile, l)
	if l > 0 {
		for i, v := range resp.List {
			users[i] = &model.UserProfile{
				Id:        v.Id,
				Username:  v.Username,
				AvatarUrl: v.AvatarUrl,
			}
		}
	}

	return users, nil
}
