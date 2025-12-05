package rpc

import (
	"context"
	build "myreel/app/video/controller/rpc/pack"
	"myreel/app/video/domain/model"
	"myreel/app/video/usecase"
	video "myreel/kitex_gen/video"
	base "myreel/pkg/base/context"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct {
	useCase usecase.VideoUseCase
}

func NewVideoServiceImpl(uc usecase.VideoUseCase) *VideoServiceImpl {
	return &VideoServiceImpl{useCase: uc}
}

// VideoStream implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoStream(ctx context.Context, req *video.VideoStreamRequest) (resp *video.VideoStreamResponse, err error) {
	resp = new(video.VideoStreamResponse)

	videos, err := s.useCase.GetVideosByLatestTime(ctx, req.LatestTime)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}

	resp.Base = base.BuildSuccessResp()
	resp.Data = build.BuildVideoList(build.BuildVideos(videos), &video.Pagination{Total: int64(len(videos))})
	return
}

// Publish implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoUploadToken(ctx context.Context, req *video.GetVideoUploadTokenRequest) (resp *video.GetVideoUploadTokenResponse, err error) {
	resp = new(video.GetVideoUploadTokenResponse)

	token, err := s.useCase.GetVideoUplaodToken(ctx, req.Suffix, req.UserId)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}

	resp.Base = base.BuildSuccessResp()
	resp.Data = build.BuildUpyunToken(token)
	return
}

// GetVideoCoverUploadToken implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoCoverUploadToken(ctx context.Context, req *video.GetVideoCoverUploadTokenRequest) (resp *video.GetVideoCoverUploadTokenResponse, err error) {
	resp = new(video.GetVideoCoverUploadTokenResponse)

	token, err := s.useCase.GetVideoCoverUplaodToken(ctx, req.Suffix, req.UserId)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}

	resp.Base = base.BuildSuccessResp()
	resp.Data = build.BuildUpyunToken(token)
	return
}

// SaveVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) SaveVideo(ctx context.Context, req *video.SaveVideoRequest) (resp *video.SaveVideoResponse, err error) {
	resp = new(video.SaveVideoResponse)

	err = s.useCase.SaveVideo(ctx, &model.Video{
		Uid:         req.UserId,
		Title:       req.Title,
		Description: req.Description,
		CoverUrl:    req.CoverUrl,
		VideoUrl:    req.VideoUrl,
	})
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}

	resp.Base = base.BuildSuccessResp()
	return
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	resp = new(video.PublishListResponse)

	videos, pagination, err := s.useCase.GetVideosByUserId(ctx, req.Uid, req.Cursor, req.Limit)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}

	resp.Base = base.BuildSuccessResp()
	resp.Data = build.BuildVideoList(build.BuildVideos(videos), build.BuildPagination(pagination))
	return
}

// Popular implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Popular(ctx context.Context, req *video.PopularRequest) (resp *video.PopularResponse, err error) {
	resp = new(video.PopularResponse)

	videos, pagination, err := s.useCase.GetVideosGroupByVisitCount(ctx, req.Cursor, req.Limit)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}

	resp.Base = base.BuildSuccessResp()
	resp.Data = build.BuildVideoList(build.BuildVideos(videos), build.BuildPagination(pagination))
	return
}

// Search implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Search(ctx context.Context, req *video.SearchRequest) (resp *video.SearchResponse, err error) {
	resp = new(video.SearchResponse)

	videos, pagination, err := s.useCase.GetVideosByKeywords(ctx, req.Keywords, req.Username, req.FromDate, req.ToDate, req.Cursor, req.Limit)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}

	resp.Base = base.BuildSuccessResp()
	resp.Data = build.BuildVideoList(build.BuildVideos(videos), build.BuildPagination(pagination))
	return
}

// VideoLikeAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoLikeAction(ctx context.Context, req *video.VideoLikeActionRequest) (resp *video.VideoLikeActionResponse, err error) {
	resp = new(video.VideoLikeActionResponse)

	err = s.useCase.VideoLikeAction(ctx, req.VideoId, req.ActionType)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}

	resp.Base = base.BuildSuccessResp()
	return
}

// GetVideosByIds implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideosByIds(ctx context.Context, req *video.GetVideosByIdsRequest) (resp *video.GetVideosByIdsResponse, err error) {
	resp = new(video.GetVideosByIdsResponse)

	videos, err := s.useCase.GetVideosByIds(ctx, req.Id)
	if err != nil {
		resp.Base = base.BuildBaseResp(err)
		return
	}

	resp.Base = base.BuildSuccessResp()
	resp.Items = build.BuildVideos(videos)
	return
}
