package pack

import (
	"encoding/json"
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

	id, err := parseInt64("Id")
	if err != nil {
		return nil, errno.Errorf(errno.InternalRedisErrorCode, "parse id: %v", err)
	}

	uid, err := parseInt64("Uid")
	if err != nil {
		return nil, errno.Errorf(errno.InternalRedisErrorCode, "parse uid: %v", err)
	}

	visitCount, err := parseInt64("VisitCount")
	if err != nil {
		return nil, errno.Errorf(errno.InternalRedisErrorCode, "parse visit_count: %v", err)
	}

	likeCount, err := parseInt64("LikeCount")
	if err != nil {
		return nil, errno.Errorf(errno.InternalRedisErrorCode, "parse like_count: %v", err)
	}

	commentCount, err := parseInt64("CommentCount")
	if err != nil {
		return nil, errno.Errorf(errno.InternalRedisErrorCode, "parse comment_count: %v", err)
	}

	createdAt, err := parseInt64("CreatedAt")
	if err != nil {
		return nil, errno.Errorf(errno.InternalRedisErrorCode, "parse created_at: %v", err)
	}

	updatedAt, err := parseInt64("UpdatedAt")
	if err != nil {
		return nil, errno.Errorf(errno.InternalRedisErrorCode, "parse updated_at: %v", err)
	}

	return &model.Video{
		Id:           id,
		Uid:          uid,
		Title:        data["Title"],
		Description:  data["Description"],
		VideoUrl:     data["VideoUrl"],
		CoverUrl:     data["CoverUrl"],
		VisitCount:   visitCount,
		LikeCount:    likeCount,
		CommentCount: commentCount,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}, nil
}

func VideoToMap(v *model.Video) (map[string]interface{}, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, errno.Errorf(errno.InternalRedisErrorCode, "redis: failed to marshal video: %v", err)
	}

	var fields map[string]interface{}
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, errno.Errorf(errno.InternalRedisErrorCode, "redis: failed to unmarshal map: %v", err)
	}

	return fields, nil
}
