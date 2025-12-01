package mysql

import (
	"context"
	"errors"
	"myreel/app/user/domain/model"
	"myreel/app/user/domain/repository"
	"myreel/pkg/errno"

	"gorm.io/gorm"
)

type userDB struct {
	client *gorm.DB
}

func NewUserDB(client *gorm.DB) repository.UserDB {
	return &userDB{client: client}
}

func (db *userDB) Magrate() error {
	if err := db.client.AutoMigrate(&User{}); err != nil {
		return errno.NewErrNo(errno.InternalDatabaseErrorCode, "mysql: failed to auto magrate user model")
	}
	return nil
}

func (db *userDB) CreateUser(ctx context.Context, user *model.User) error {
	u := &User{
		Id:        user.Id,
		Username:  user.Username,
		Password:  user.Password,
		AvatarUrl: user.AvatarUrl,
	}
	if err := db.client.WithContext(ctx).Create(u).Error; err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to create user: %v", err)
	}
	return nil
}

func (db *userDB) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user User
	err := db.client.WithContext(ctx).Where("username = ?", username).
		Where("deleted_at IS NULL").
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.UserNotFound
		}
		return nil, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get user by username: %v", err)
	}
	return &model.User{
		Id:        user.Id,
		Username:  user.Username,
		Password:  user.Password,
		AvatarUrl: user.AvatarUrl,
	}, nil
}

func (db *userDB) GetUserById(ctx context.Context, id int64) (*model.User, error) {
	var user User
	if err := db.client.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get user by id: %v", err)
	}
	return &model.User{
		Id:        user.Id,
		Username:  user.Username,
		Password:  user.Password,
		AvatarUrl: user.AvatarUrl,
	}, nil
}

func (db *userDB) SetAvatar(ctx context.Context, id string, url string) error {
	if err := db.client.WithContext(ctx).Model(&User{}).Where("id = ?", id).Update("avatar_url", url).Error; err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to set avatar: %v", err)
	}
	return nil
}
