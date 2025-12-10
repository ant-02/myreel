package rpc

import (
	"context"
	"myreel/kitex_gen/follow"
	"myreel/pkg/base/client"
	"myreel/pkg/errno"
	"myreel/pkg/logger"
	"myreel/pkg/util"

	api "myreel/app/gateway/model/api/follow"
)

func InitFollowClient() {
	c, err := client.InitFollowRPC()
	if err != nil {
		logger.Fatalf("api.rpc.like InitFollowRpc failed, err is %v", err)
	}
	followClient = *c
}

func FollowActionRPC(ctx context.Context, req *follow.FollowActionRequest) error {
	resp, err := followClient.FollowAction(ctx, req)
	if err != nil {
		logger.Errorf("FollowActionRPC: RPC called failed: %v", err.Error())
		return errno.InternalServiceError.WithError(err)
	}
	if !util.IsSuccess(resp.Base) {
		return errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return nil
}

func GetFolloweringsByIdsRPC(ctx context.Context, req *follow.FolloweringListRequest) (*api.FolloweringListResponse, error) {
	resp, err := followClient.FolloweringList(ctx, req)
	if err != nil {
		logger.Errorf("GetFolloweringsByIdsRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !util.IsSuccess(resp.Base) {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}

	l := len(resp.Data.Items)
	user := make([]*api.User, l)
	if l > 0 {
		for i, v := range resp.Data.Items {
			user[i] = &api.User{
				Id:        v.Id,
				Username:  v.Username,
				AvatarUrl: v.AvatarUrl,
			}
		}
	}

	return &api.FolloweringListResponse{
		Items: user,
		Pagination: &api.Pagination{
			NextCursor: resp.Data.Pagination.NextCursor,
			PrevCursor: resp.Data.Pagination.PrevCursor,
			Total:      resp.Data.Pagination.Total,
		},
	}, nil
}

func GetFolloweredsByIdsRPC(ctx context.Context, req *follow.FolloweredListRequest) (*api.FolloweredListResponse, error) {
	resp, err := followClient.FolloweredList(ctx, req)
	if err != nil {
		logger.Errorf("GetFolloweredsByIdsRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !util.IsSuccess(resp.Base) {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}

	l := len(resp.Data.Items)
	user := make([]*api.User, l)
	if l > 0 {
		for i, v := range resp.Data.Items {
			user[i] = &api.User{
				Id:        v.Id,
				Username:  v.Username,
				AvatarUrl: v.AvatarUrl,
			}
		}
	}

	return &api.FolloweredListResponse{
		Items: user,
		Pagination: &api.Pagination{
			NextCursor: resp.Data.Pagination.NextCursor,
			PrevCursor: resp.Data.Pagination.PrevCursor,
			Total:      resp.Data.Pagination.Total,
		},
	}, nil
}

func GetFriendsByIdRPC(ctx context.Context, req *follow.FriendListRequest) (*api.FriendListResponse, error) {
	resp, err := followClient.FriendList(ctx, req)
	if err != nil {
		logger.Errorf("GetFriendsByIdRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !util.IsSuccess(resp.Base) {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}

	l := len(resp.Data.Items)
	user := make([]*api.User, l)
	if l > 0 {
		for i, v := range resp.Data.Items {
			user[i] = &api.User{
				Id:        v.Id,
				Username:  v.Username,
				AvatarUrl: v.AvatarUrl,
			}
		}
	}

	return &api.FriendListResponse{
		Items: user,
		Pagination: &api.Pagination{
			NextCursor: resp.Data.Pagination.NextCursor,
			PrevCursor: resp.Data.Pagination.PrevCursor,
			Total:      resp.Data.Pagination.Total,
		},
	}, nil
}
