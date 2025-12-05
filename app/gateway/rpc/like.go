package rpc

import (
	"context"
	api "myreel/app/gateway/model/api/like"
	"myreel/app/gateway/model/api/video"
	"myreel/kitex_gen/like"
	"myreel/pkg/base/client"
	"myreel/pkg/errno"
	"myreel/pkg/logger"
	"myreel/pkg/util"
)

func InitLikeClient() {
	c, err := client.InitLikeRPC()
	if err != nil {
		logger.Fatalf("api.rpc.like InitLikeRPC failed, err is %v", err)
	}
	likeClient = *c
}

func LikeActionRPC(ctx context.Context, req *like.LikeActionRequest) error {
	resp, err := likeClient.LikeAction(ctx, req)
	if err != nil {
		logger.Errorf("LikeActionRPC: RPC called failed: %v", err.Error())
		return errno.InternalServiceError.WithError(err)
	}
	if !util.IsSuccess(resp.Base) {
		return errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return nil
}

func LikeListRPC(ctx context.Context, req *like.LikeListRequest) (*api.LikeListResponse, error) {
	resp, err := likeClient.LikeList(ctx, req)
	if err != nil {
		logger.Errorf("LikeListRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !util.IsSuccess(resp.Base) {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}

	l := len(resp.Data.Items)
	videos := make([]*video.Video, l)
	if l > 0 {
		for i, v := range resp.Data.Items {
			videos[i] = &video.Video{
				Id:           v.Id,
				Uid:          v.Uid,
				Title:        v.Title,
				Description:  v.Description,
				VideoUrl:     v.VideoUrl,
				CoverUrl:     v.CoverUrl,
				VisitCount:   v.VisitCount,
				LikeCount:    v.LikeCount,
				CommentCount: v.CommentCount,
			}
		}
	}

	return &api.LikeListResponse{
		Items: videos,
		Pagination: &video.Pagination{
			NextCursor: resp.Data.Pagination.NextCursor,
			PrevCursor: resp.Data.Pagination.PrevCursor,
			Total:      resp.Data.Pagination.Total,
		},
	}, nil
}
