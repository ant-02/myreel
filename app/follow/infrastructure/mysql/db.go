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
	if err := db.client.AutoMigrate(&Follow{}, &Group{}, &GroupMember{}); err != nil {
		return errno.NewErrNo(errno.InternalDatabaseErrorCode, "mysql: failed to auto magrate follow group groupMember model")
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

func (db *followDB) GetFriendIdsById(ctx context.Context, id, limit int64, cursor time.Time) ([]*model.FolloweredIdWithTime, int64, error) {
	var fs []*model.FolloweredIdWithTime
	var total int64

	tx := db.client.WithContext(ctx)
	if err := tx.Raw(`
		SELECT COUNT(*)
		FROM follows f1
		WHERE f1.followering_id = ? 
			AND f1.status = 1
			AND EXISTS (
				SELECT 1
				FROM follows f2
				WHERE f2.followering_id = f1.followered_id
					AND f2.status = 1
					AND f2.followered_id = ?
			)
	`, id, id).
		Scan(&total).
		Error; err != nil {
		return nil, 0, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get friends count by id: %v", err)
	}

	if err := tx.Raw(`
		SELECT f1.followered_id, f1.created_at
		FROM follows f1
		WHERE f1.followering_id = ? 
			AND	f1.created_at < ?
			AND f1.status = 1
			AND EXISTS (
				SELECT 1
				FROM follows f2
				WHERE f2.followering_id = f1.followered_id
					AND f2.status = 1
					AND f2.followered_id = ?
			)
		ORDER BY f1.created_at DESC
		LIMIT ?
	`, id, cursor, id, limit).
		Scan(&fs).
		Error; err != nil {
		return nil, 0, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get friends by id: %v", err)
	}

	return fs, total, nil
}

func (db *followDB) CreateGroup(ctx context.Context, g *model.Group) error {
	group := &Group{
		Id:        g.Id,
		Name:      g.Name,
		CreatorId: g.CreatorId,
	}
	if err := db.client.WithContext(ctx).Create(&group).Error; err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to create group: %v", err)
	}
	return nil
}

func (db *followDB) CreateGroupMember(ctx context.Context, gm *model.GroupMember) error {
	groupMember := &GroupMember{
		Id:      gm.Id,
		GroupId: gm.GroupId,
		UserId:  gm.UserId,
	}
	if err := db.client.WithContext(ctx).Create(&groupMember).Error; err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to create group member: %v", err)
	}
	return nil
}

func (db *followDB) GetGroupIdsByJoined(ctx context.Context, userId int64) ([]int64, error) {
	var ids []int64
	if err := db.client.WithContext(ctx).
		Model(&GroupMember{}).
		Where("user_id = ? ", userId).
		Select("group_id").
		Find(&ids).
		Error; err != nil {
		return nil, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get group id list by joined: %v", err)
	}
	return ids, nil
}

func (db *followDB) GetGroupById(ctx context.Context, id int64) (*model.Group, error) {
	var g Group
	if err := db.client.WithContext(ctx).
		First(&g, id).
		Error; err != nil {
		return nil, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get group by id: %v", err)
	}
	return &model.Group{
		Id:        g.Id,
		Name:      g.Name,
		CreatorId: g.CreatorId,
		CreatedAt: g.CreatedAt.Unix(),
	}, nil
}

func (db *followDB) GetGroupByCreator(ctx context.Context, creatorId int64) ([]*model.Group, error) {
	var gs []*Group
	if err := db.client.WithContext(ctx).
		Where("creator_id = ?", creatorId).
		Find(&gs).
		Error; err != nil {
		return nil, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get group list by creator: %v", err)
	}
	l := len(gs)
	result := make([]*model.Group, l)
	if l > 0 {
		for i, v := range gs {
			result[i] = &model.Group{
				Id:        v.Id,
				Name:      v.Name,
				CreatorId: v.CreatorId,
				CreatedAt: v.CreatedAt.Unix(),
			}
		}
	}
	return result, nil
}
