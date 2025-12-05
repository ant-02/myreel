package pack

import (
	"myreel/app/video/domain/model"
	"myreel/pkg/errno"
	"strconv"
)

func MapToVideo(data map[string]string) (*model.Video, error) {
	// 辅助函数：安全解析 int64
	parseInt64 := func(key string) (int64, error) {
		if s, ok := data[key]; ok && s != "" {
			return strconv.ParseInt(s, 10, 64)
		}
		return 0, errno.Errorf(errno.InternalRedisErrorCode, "missing or empty field: %s", key)
	}

	id, err := parseInt64("id")
	if err != nil {
		return nil, errno.Errorf(errno.InternalRedisErrorCode, "parse id: %v", err)
	}

	uid, err := parseInt64("uid")
	if err != nil {
		return nil, errno.Errorf(errno.InternalRedisErrorCode, "parse uid: %v", err)
	}

	visitCount, err := parseInt64("visit_count")
	if err != nil {
		return nil, errno.Errorf(errno.InternalRedisErrorCode, "parse visit_count: %v", err)
	}

	likeCount, err := parseInt64("like_count")
	if err != nil {
		return nil, errno.Errorf(errno.InternalRedisErrorCode, "parse like_count: %v", err)
	}

	commentCount, err := parseInt64("comment_count")
	if err != nil {
		return nil, errno.Errorf(errno.InternalRedisErrorCode, "parse comment_count: %v", err)
	}

	createdAt, err := parseInt64("created_at")
	if err != nil {
		return nil, errno.Errorf(errno.InternalRedisErrorCode, "parse created_at: %v", err)
	}

	updatedAt, err := parseInt64("updated_at")
	if err != nil {
		return nil, errno.Errorf(errno.InternalRedisErrorCode, "parse updated_at: %v", err)
	}

	return &model.Video{
		Id:           id,
		Uid:          uid,
		Title:        data["title"],
		Description:  data["description"],
		VideoUrl:     data["video_url"],
		CoverUrl:     data["cover_url"],
		VisitCount:   visitCount,
		LikeCount:    likeCount,
		CommentCount: commentCount,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}, nil
}
