package mysql

import (
	"context"
	"errors"
	"myreel/app/follow/domain/model"
	"myreel/app/follow/domain/repository"
	"myreel/pkg/errno"

	"gorm.io/gorm"
)

type followDB struct {
	client *gorm.DB
}

func NewFollowDB(client *gorm.DB) repository.FollowDB {
	return &followDB{client: client}
}

func (db *followDB) Magrate() error {
	if err := db.client.AutoMigrate(&Follow{}); err != nil {
		return errno.NewErrNo(errno.InternalDatabaseErrorCode, "mysql: failed to auto magrate follow model")
	}
	return nil
}

func (db *followDB) GetFollowByUserIdAndToUserId(ctx context.Context, userId, toUserId int64) (*model.Follow, error) {
	var follow Follow
	if err := db.client.WithContext(ctx).
		Model(&Follow{}).
		Where("followering_id = ? and followered_id = ?", userId, toUserId).
		First(&follow).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.FollowNotFound
		}
		return nil, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get follow by user id and to User id: %v", err)
	}
	return &model.Follow{
		Id:            follow.Id,
		FolloweringId: follow.FolloweringId,
		FolloweredId:  follow.FolloweredId,
		Status:        follow.Status,
		CreatedAt:     follow.CreatedAt.Unix(),
	}, nil
}

func (db *followDB) SetFollowStatus(ctx context.Context, id, status int64) error {
	if err := db.client.WithContext(ctx).
		Model(&Follow{}).
		Where("id = ?", id).
		Update("status", status).
		Error; err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, " mysql: failed to set follow status: %v", err)
	}
	return nil
}

func (db *followDB) CreateFollow(ctx context.Context, f *model.Follow) error {
	follow := &Follow{
		Id:            f.Id,
		FolloweringId: f.FolloweringId,
		FolloweredId:  f.FolloweredId,
		Status:        1,
	}
	if err := db.client.WithContext(ctx).
		Create(&follow).
		Error; err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to create follow: %v", err)
	}
	return nil
}
