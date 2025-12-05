package rpc

import (
	"context"
	"myreel/app/comment/usecase"
	comment "myreel/kitex_gen/comment"
	base "myreel/pkg/base/context"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct {
	useCase usecase.CommentUseCase
}

func NewCommentServiceImpl(uc usecase.CommentUseCase) *CommentServiceImpl {
	return &CommentServiceImpl{useCase: uc}
}

// CommentPublish implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentPublish(ctx context.Context, req *comment.CommentPublishRequest) (resp *comment.CommentPublishResponse, err error) {
	resp = new(comment.CommentPublishResponse)

	err = s.useCase.CommentPublish(ctx, req.VideoId, req.CommentId, req.UserId, req.Content)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}

	resp.Base = base.BuildSuccessResp()
	return
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}

// Delete implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) Delete(ctx context.Context, req *comment.DeleteRequest) (resp *comment.DeleteResponse, err error) {
	// TODO: Your code here...
	return
}
