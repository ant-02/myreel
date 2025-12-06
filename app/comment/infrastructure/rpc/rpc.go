package rpc

import (
	"context"
	"myreel/app/comment/domain/repository"
	"myreel/kitex_gen/video"
	"myreel/kitex_gen/video/videoservice"
	"myreel/pkg/errno"
)

type commentRpcImpl struct {
	video videoservice.Client
}

func NewCommentRpcImpl(v videoservice.Client) repository.RpcPort {
	return &commentRpcImpl{
		video: v,
	}
}

func (rpc *commentRpcImpl) AddCommentCount(ctx context.Context, id int64) error {
	resp, err := rpc.video.AddCommentCount(ctx, &video.AddCommentCountRequest{
		Id: id,
	})
	if err != nil {
		return errno.Errorf(errno.InternalRPCErrorCode, "rpc: failed to add video comment count: %v", err)
	}

	if resp.Base.Code != errno.SuccessCode {
		return errno.NewErrNo(errno.InternalRPCErrorCode, "rpc: failed to add video comment count")
	}

	return nil
}
