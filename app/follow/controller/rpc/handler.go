package rpc

import (
	"context"
	"myreel/app/follow/usecase"
	follow "myreel/kitex_gen/follow"
	base "myreel/pkg/base/context"
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
	// TODO: Your code here...
	return
}

// FolloweredList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FolloweredList(ctx context.Context, req *follow.FolloweredListRequest) (resp *follow.FolloweredListResponse, err error) {
	// TODO: Your code here...
	return
}

// FriendList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FriendList(ctx context.Context, req *follow.FriendListRequest) (resp *follow.FriendListResponse, err error) {
	// TODO: Your code here...
	return
}
