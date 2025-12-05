package rpc

import (
	"context"
	"myreel/app/like/domain/repository"
	"myreel/kitex_gen/video"
	"myreel/kitex_gen/video/videoservice"
	"myreel/pkg/errno"
)

type likeRpcImpl struct {
	video videoservice.Client
}

func NewVideoRpcImpl(v videoservice.Client) repository.RpcPort {
	return &likeRpcImpl{
		video: v,
	}
}

func (rpc *likeRpcImpl) VideoLikeAction(ctx context.Context, videoId, actionType int64) error {
	resp, err := rpc.video.VideoLikeAction(ctx, &video.VideoLikeActionRequest{
		VideoId: videoId,
		ActionType: actionType,
	})
	if err != nil {
		return errno.Errorf(errno.InternalRPCErrorCode, "rpc: failed to action video likes: %v", err)
	}

	if resp.Base.Code != errno.SuccessCode {
		return errno.NewErrNo(errno.InternalRPCErrorCode, "rpc: ffailed to action video likes")
	}

	return nil
}
