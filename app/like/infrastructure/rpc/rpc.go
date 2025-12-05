package rpc

import (
	"context"
	"myreel/app/like/domain/model"
	"myreel/app/like/domain/repository"
	"myreel/kitex_gen/video"
	"myreel/kitex_gen/video/videoservice"
	"myreel/pkg/errno"
)

type likeRpcImpl struct {
	video videoservice.Client
}

func NewLikeRpcImpl(v videoservice.Client) repository.RpcPort {
	return &likeRpcImpl{
		video: v,
	}
}

func (rpc *likeRpcImpl) VideoLikeAction(ctx context.Context, videoId, actionType int64) error {
	resp, err := rpc.video.VideoLikeAction(ctx, &video.VideoLikeActionRequest{
		VideoId:    videoId,
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

func (rpc *likeRpcImpl) GetVideosByIds(ctx context.Context, ids []int64) ([]*model.Video, error) {
	resp, err := rpc.video.GetVideosByIds(ctx, &video.GetVideosByIdsRequest{
		Id: ids,
	})
	if err != nil {
		return nil, errno.Errorf(errno.InternalRPCErrorCode, "rpc: failed to action video likes: %v", err)
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(errno.InternalRPCErrorCode, "rpc: ffailed to action video likes")
	}

	l := len(resp.Items)
	videos := make([]*model.Video, l)
	if l > 0 {
		for i, item := range resp.Items {
			videos[i] = &model.Video{
				Id:           item.Id,
				Uid:          item.Uid,
				Title:        item.Title,
				Description:  item.Description,
				VideoUrl:     item.VideoUrl,
				CoverUrl:     item.CoverUrl,
				LikeCount:    *item.LikeCount,
				CommentCount: *item.CommentCount,
				VisitCount:   *item.VisitCount,
				CreatedAt:    item.CreatedAt,
				UpdatedAt:    item.UpdatedAt,
			}
		}
	}
	return videos, nil
}
