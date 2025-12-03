package pack

import (
	domainModel "myreel/app/video/domain/model"
	"myreel/kitex_gen/model"
	"myreel/kitex_gen/video"
	"myreel/pkg/upyun"
)

func BuildVideo(v *domainModel.Video) *video.Video {
	if v == nil {
		return nil
	}

	vc, lc, cc := v.VisitCount, v.LikeCount, v.CommentCount
	return &video.Video{
		Id:           v.Id,
		Uid:          v.Uid,
		Title:        v.Title,
		Description:  v.Description,
		VideoUrl:     v.VideoUrl,
		CoverUrl:     v.CoverUrl,
		VisitCount:   &vc,
		LikeCount:    &lc,
		CommentCount: &cc,
	}
}

func BuildVideos(vs []*domainModel.Video) []*video.Video {
	l := len(vs)
	if l == 0 {
		return nil
	}

	videos := make([]*video.Video, 0, l)
	for i, v := range vs {
		videos[i] = BuildVideo(v)
	}
	return videos
}

func BuildVideoList(vs []*video.Video, t int64) *video.VideoList {
	return &video.VideoList{
		Items: vs,
		Total: &t,
	}
}

func BuildUpyunToken(t *upyun.UpyunToken) *model.UpyunToken {
	return &model.UpyunToken{
		Policy:        t.Policy,
		Authorization: t.Authorization,
		Bucket:        t.Bucket,
	}
}
