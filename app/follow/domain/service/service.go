package service

import (
	"context"
	"myreel/app/follow/domain/model"
	"myreel/pkg/errno"
	"time"
)

func (fs *followService) FollowAction(ctx context.Context, userId, toUserId, actionType int64) error {
	follow, err := fs.db.GetFollowByUserIdAndToUserId(ctx, userId, toUserId)
	if err != nil {
		return err
	}
	if actionType == follow.Status {
		return nil
	}
	err = fs.db.SetFollowStatus(ctx, follow.Id, actionType)
	if err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to set follow status")
	}
	return nil
}

func (fs *followService) GenerateFollowId() (int64, error) {
	id, err := fs.sf.Generate()
	if err != nil {
		return 0, err
	}
	return id, err
}

func (fs *followService) CreateFollow(ctx context.Context, f *model.Follow) error {
	if err := fs.db.CreateFollow(ctx, f); err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to create follow").WithError(err)
	}
	return nil
}

func (fs *followService) GetUsersByFolloweredId(ctx context.Context, userId, cursor, limit int64) ([]*model.UserProfile, *model.Pagination, error) {
	t := time.Now()
	if cursor > 0 {
		t = time.Unix(cursor, 10)
	}
	fts, total, err := fs.db.GetUserIdsByFolloweredId(ctx, userId, limit, t)
	if err != nil {
		return nil, nil, errno.Errorf(errno.InternalServiceErrorCode, "failed to get follows by user id: %v", err)
	}

	nextCursor := cursor
	l := len(fts)
	ids := make([]int64, l)
	if l > 0 {
		nextCursor = fts[l-1].CreatedAt.Unix()
		for i, v := range fts {
			ids[i] = v.FolloweringId
		}
	}

	users, err := fs.rpc.GetUsersByIdsRPC(ctx, ids)
	if err != nil {
		return nil, nil, errno.Errorf(errno.InternalServiceErrorCode, "failed to get users by ids: %v", err)
	}

	return users, &model.Pagination{
		NextCursor: nextCursor,
		PrevCursor: cursor,
		Total:      total,
	}, nil
}

func (fs *followService) GetUsersByFolloweringId(ctx context.Context, userId, cursor, limit int64) ([]*model.UserProfile, *model.Pagination, error) {
	t := time.Now()
	if cursor > 0 {
		t = time.Unix(cursor, 10)
	}
	fts, total, err := fs.db.GetUserIdsByFolloweringId(ctx, userId, limit, t)
	if err != nil {
		return nil, nil, errno.Errorf(errno.InternalServiceErrorCode, "failed to get follows by user id: %v", err)
	}

	nextCursor := cursor
	l := len(fts)
	ids := make([]int64, l)
	if l > 0 {
		nextCursor = fts[l-1].CreatedAt.Unix()
		for i, v := range fts {
			ids[i] = v.FolloweredId
		}
	}

	users, err := fs.rpc.GetUsersByIdsRPC(ctx, ids)
	if err != nil {
		return nil, nil, errno.Errorf(errno.InternalServiceErrorCode, "failed to get users by ids: %v", err)
	}

	return users, &model.Pagination{
		NextCursor: nextCursor,
		PrevCursor: cursor,
		Total:      total,
	}, nil
}

func (fs *followService) GetFriendsById(ctx context.Context, id, cursor, limit int64) ([]*model.UserProfile, *model.Pagination, error) {
	t := time.Now()
	if cursor > 0 {
		t = time.Unix(cursor, 10)
	}
	fts, total, err := fs.db.GetFriendIdsById(ctx, id, limit, t)
	if err != nil {
		return nil, nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get friends by id").WithError(err)
	}

	nextCursor := cursor
	l := len(fts)
	ids := make([]int64, l)
	if l > 0 {
		nextCursor = fts[l-1].CreatedAt.Unix()
		for i, v := range fts {
			ids[i] = v.FolloweredId
		}
	}

	users, err := fs.rpc.GetUsersByIdsRPC(ctx, ids)
	if err != nil {
		return nil, nil, errno.Errorf(errno.InternalServiceErrorCode, "failed to get users by ids: %v", err)
	}

	return users, &model.Pagination{
		NextCursor: nextCursor,
		PrevCursor: cursor,
		Total:      total,
	}, nil
}

func (fs *followService) GenerateGroupId() (int64, error) {
	id, err := fs.sf.Generate()
	if err != nil {
		return 0, err
	}
	return id, err
}

func (fs *followService) CreateGroup(ctx context.Context, g *model.Group) error {
	if err := fs.db.CreateGroup(ctx, g); err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to create group").WithError(err)
	}
	return nil
}

func (fs *followService) GenerateGroupMemberId() (int64, error) {
	id, err := fs.sf.Generate()
	if err != nil {
		return 0, err
	}
	return id, err
}

func (fs *followService) CreateGroupMember(ctx context.Context, gm *model.GroupMember) error {
	if err := fs.db.CreateGroupMember(ctx, gm); err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to create group member").WithError(err)
	}
	return nil
}

func (fs *followService) GetGroupByJoined(ctx context.Context, userId int64) ([]*model.Group, error) {
	ids, err := fs.db.GetGroupIdsByJoined(ctx, userId)
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get group id list by joined").WithError(err)
	}

	l := len(ids)
	gs := make([]*model.Group, l)
	if l > 0 {
		for i, v := range ids {
			g, err := fs.db.GetGroupById(ctx, v)
			if err != nil {
				return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get group by id").WithError(err)
			}
			gs[i] = g
		}
	}

	return gs, nil
}

func (fs *followService) GetGroupByCreator(ctx context.Context, creatorId int64) ([]*model.Group, error) {
	gs, err := fs.db.GetGroupByCreator(ctx, creatorId)
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get group list by creator").WithError(err)
	}
	return gs, nil
}
