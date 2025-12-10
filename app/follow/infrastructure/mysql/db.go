package mysql

import (
	"context"
	"errors"
	"myreel/app/follow/domain/model"
	"myreel/app/follow/domain/repository"
	"myreel/pkg/errno"
	"time"

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

func (db *followDB) GetUserIdsByFolloweredId(ctx context.Context, userId, limit int64, cursor time.Time) ([]*model.FolloweringIdWithTime, int64, error) {
	var fs []*model.FolloweringIdWithTime
	var total int64
	tx := db.client.WithContext(ctx).
		Model(&Follow{}).
		Where("followered_id = ? and status = 1", userId)

	err := tx.Count(&total).Error
	if err != nil {
		return nil, 0, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get follow count by followered id: %v", err)
	}

	err = tx.Where("created_at < ?", cursor).
		Limit(int(limit)).
		Order("created_at DESC").
		Select("followering_id", "created_at").
		Find(&fs).
		Error
	if err != nil {
		return nil, 0, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get follows by followered id: %v", err)
	}

	return fs, total, nil
}

func (db *followDB) GetUserIdsByFolloweringId(ctx context.Context, userId, limit int64, cursor time.Time) ([]*model.FolloweredIdWithTime, int64, error) {
	var fs []*model.FolloweredIdWithTime
	var total int64
	tx := db.client.WithContext(ctx).
		Model(&Follow{}).
		Where("followering_id = ? and status = 1", userId)

	err := tx.Count(&total).Error
	if err != nil {
		return nil, 0, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get follow count by followering id: %v", err)
	}

	err = tx.Where("created_at < ?", cursor).
		Limit(int(limit)).
		Order("created_at DESC").
		Select("followered_id", "created_at").
		Find(&fs).
		Error
	if err != nil {
		return nil, 0, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get follows by followering id: %v", err)
	}

	return fs, total, nil
}
