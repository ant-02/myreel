package rpc

import (
	"context"
	api "myreel/app/gateway/model/api/video"
	"myreel/kitex_gen/video"
	"myreel/pkg/base/client"
	"myreel/pkg/errno"
	"myreel/pkg/logger"
	"myreel/pkg/util"
)

func InitVideoClient() {
	c, err := client.InitVideoRPC()
	if err != nil {
		logger.Fatalf("api.rpc.video InitVideoRPC failed, err is %v", err)
	}
	videoClient = *c
}

func VideoSteamRPC(ctx context.Context, req *video.VideoStreamRequest) (*api.VideoStreamResponse, error) {
	resp, err := videoClient.VideoStream(ctx, req)
	if err != nil {
		logger.Errorf("VideoSteamRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !util.IsSuccess(resp.Base) {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}

	l := len(resp.Data.Items)
	items := make([]*api.Video, l)
	if l > 0 {
		for i, v := range resp.Data.Items {
			vc, lc, cc := *v.VisitCount, *v.LikeCount, *v.CommentCount
			items[i] = &api.Video{
				Id:           v.Id,
				Uid:          v.Uid,
				CoverUrl:     v.CoverUrl,
				VideoUrl:     v.VideoUrl,
				Title:        v.Title,
				Description:  v.Description,
				VisitCount:   &vc,
				LikeCount:    &lc,
				CommentCount: &cc,
			}
		}
	}

	return &api.VideoStreamResponse{
		Items: items,
		Total: &resp.Data.Pagination.Total,
	}, nil
}

func GetVideoUploadTokenRPC(ctx context.Context, req *video.GetVideoUploadTokenRequest) (*api.GetVideoUploadTokenResponse, error) {
	resp, err := videoClient.GetVideoUploadToken(ctx, req)
	if err != nil {
		logger.Errorf("GetVideoUploadTokenRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !util.IsSuccess(resp.Base) {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}

	return &api.GetVideoUploadTokenResponse{
		Policy:        resp.Data.Policy,
		Authorization: resp.Data.Authorization,
		Bucket:        resp.Data.Bucket,
	}, nil
}

func GetVideoCoverUploadTokenRPC(ctx context.Context, req *video.GetVideoCoverUploadTokenRequest) (*api.GetVideoCoverUploadTokenResponse, error) {
	resp, err := videoClient.GetVideoCoverUploadToken(ctx, req)
	if err != nil {
		logger.Errorf("GetVideoCoverUploadTokenRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !util.IsSuccess(resp.Base) {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}

	return &api.GetVideoCoverUploadTokenResponse{
		Policy:        resp.Data.Policy,
		Authorization: resp.Data.Authorization,
		Bucket:        resp.Data.Bucket,
	}, nil
}

func SaveVideoRPC(ctx context.Context, req *video.SaveVideoRequest) error {
	resp, err := videoClient.SaveVideo(ctx, req)
	if err != nil {
		logger.Errorf("SaveVideoRPC: RPC called failed: %v", err.Error())
		return errno.InternalServiceError.WithError(err)
	}
	if !util.IsSuccess(resp.Base) {
		return errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}

	return nil
}

func PublishListRPC(ctx context.Context, req *video.PublishListRequest) (*api.PublishListResponse, error) {
	resp, err := videoClient.PublishList(ctx, req)
	if err != nil {
		logger.Errorf("PublishListRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !util.IsSuccess(resp.Base) {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}

	l := len(resp.Data.Items)
	videos := make([]*api.Video, l)
	if l > 0 {
		for i, v := range resp.Data.Items {
			videos[i] = &api.Video{
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

	return &api.PublishListResponse{
		Items: videos,
		Pagination: &api.Pagination{
			NextCursor: resp.Data.Pagination.NextCursor,
			PrevCursor: resp.Data.Pagination.PrevCursor,
			Total:      resp.Data.Pagination.Total,
		},
	}, nil
}

func PopularRPC(ctx context.Context, req *video.PopularRequest) (*api.PopularResponse, error) {
	resp, err := videoClient.Popular(ctx, req)
	if err != nil {
		logger.Errorf("PublishListRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !util.IsSuccess(resp.Base) {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}

	l := len(resp.Data.Items)
	videos := make([]*api.Video, l)
	if l > 0 {
		for i, v := range resp.Data.Items {
			videos[i] = &api.Video{
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

	return &api.PopularResponse{
		Items: videos,
		Pagination: &api.Pagination{
			NextCursor: resp.Data.Pagination.NextCursor,
			PrevCursor: resp.Data.Pagination.PrevCursor,
			Total:      resp.Data.Pagination.Total,
		},
	}, nil
}

func SearchRPC(ctx context.Context, req *video.SearchRequest) (*api.SearchResponse, error) {
	resp, err := videoClient.Search(ctx, req)
	if err != nil {
		logger.Errorf("SearchRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !util.IsSuccess(resp.Base) {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}

	l := len(resp.Data.Items)
	videos := make([]*api.Video, l)
	if l > 0 {
		for i, v := range resp.Data.Items {
			videos[i] = &api.Video{
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

	return &api.SearchResponse{
		Items: videos,
		Pagination: &api.Pagination{
			NextCursor: resp.Data.Pagination.NextCursor,
			PrevCursor: resp.Data.Pagination.PrevCursor,
			Total:      resp.Data.Pagination.Total,
		},
	}, nil
}