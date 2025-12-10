package rpc

import (
	"context"
	"myreel/app/follow/usecase"
	follow "myreel/kitex_gen/follow"
	base "myreel/pkg/base/context"

	build "myreel/app/follow/controller/rpc/pack"
)

// FollowServiceImpl implements the last service interface defined in the IDL.
type FollowServiceImpl struct {
	useCase usecase.FollowUseCase
}

func NewFollowServiceImpl(u usecase.FollowUseCase) *FollowServiceImpl {
	return &FollowServiceImpl{useCase: u}
}

// FollowAction implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowAction(ctx context.Context, req *follow.FollowActionRequest) (resp *follow.FollowActionResponse, err error) {
	resp = new(follow.FollowActionResponse)

	err = s.useCase.FollowAction(ctx, req.UserId, req.ToUserId, req.ActionType)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}
	resp.Base = base.BuildSuccessResp()
	return
}

// FolloweringList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FolloweringList(ctx context.Context, req *follow.FolloweringListRequest) (resp *follow.FolloweringListResponse, err error) {
	resp = new(follow.FolloweringListResponse)

	users, pagination, err := s.useCase.GetUsersByFolloweringId(ctx, req.UserId, req.Cursor, req.Limit)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}
	resp.Base = base.BuildSuccessResp()
	resp.Data = build.BuildUserList(build.BuildUserProfiles(users), build.BuildPagination(pagination))
	return
}

// FolloweredList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FolloweredList(ctx context.Context, req *follow.FolloweredListRequest) (resp *follow.FolloweredListResponse, err error) {
	resp = new(follow.FolloweredListResponse)

	users, pagination, err := s.useCase.GetUsersByFolloweredId(ctx, req.UserId, req.Cursor, req.Limit)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}
	resp.Base = base.BuildSuccessResp()
	resp.Data = build.BuildUserList(build.BuildUserProfiles(users), build.BuildPagination(pagination))
	return
}

// FriendList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FriendList(ctx context.Context, req *follow.FriendListRequest) (resp *follow.FriendListResponse, err error) {
	// TODO: Your code here...
	return
}
