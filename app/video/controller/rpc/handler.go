package rpc

import (
	"context"
	build "myreel/app/video/controller/rpc/pack"
	"myreel/app/video/usecase"
	video "myreel/kitex_gen/video"
	base "myreel/pkg/base/context"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct {
	useCase usecase.VideoUseCase
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
	resp.Data = build.BuildVideoList(build.BuildVideos(videos), int64(len(videos)))
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

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// Popular implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Popular(ctx context.Context, req *video.PopularRequest) (resp *video.PopularResponse, err error) {
	// TODO: Your code here...
	return
}

// Search implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Search(ctx context.Context, req *video.SearchRequest) (resp *video.SearchResponse, err error) {
	// TODO: Your code here...
	return
}
