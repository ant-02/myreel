package rpc

import (
	"context"
	build "myreel/app/like/controller/rpc/pack"
	"myreel/app/like/usecase"
	like "myreel/kitex_gen/like"
	base "myreel/pkg/base/context"
)

// LikeServiceImpl implements the last service interface defined in the IDL.
type LikeServiceImpl struct {
	useCase usecase.LikeUseCase
}

func NewLikeServiceImpl(uc usecase.LikeUseCase) *LikeServiceImpl {
	return &LikeServiceImpl{useCase: uc}
}

// LikeAction implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) LikeAction(ctx context.Context, req *like.LikeActionRequest) (resp *like.LikeActionResponse, err error) {
	resp = new(like.LikeActionResponse)

	err = s.useCase.LikeAction(ctx, req.VideoId, req.CommentId, req.UserId, req.ActionType)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}

	resp.Base = base.BuildSuccessResp()
	return
}

// LikeList implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) LikeList(ctx context.Context, req *like.LikeListRequest) (resp *like.LikeListResponse, err error) {
	resp = new(like.LikeListResponse)

	videos, pagination, err := s.useCase.GetVideosByUserLike(ctx, req.Uid, req.Cursor, req.Limit)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}

	resp.Base = base.BuildSuccessResp()
	resp.Data = build.BuildVideoList(build.BuildVideos(videos), build.BuildPagination(pagination))
	return
}
