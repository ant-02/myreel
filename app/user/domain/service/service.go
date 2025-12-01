package service

import (
	"context"
	"myreel/app/user/domain/model"
	"myreel/pkg/constants"
	"myreel/pkg/errno"
	"myreel/pkg/util"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (us *userService) GenerateUserId() (int64, error) {
	id, err := us.sf.Generate()
	if err != nil {
		return 0, err
	}
	return id, err
}

func (us *userService) EncryptPassword(pwd string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(pwd), constants.UserDefaultEncryptPasswordCost)
	if err != nil {
		return "", err
	}
	return string(password), nil
}

func (us *userService) UserRegister(ctx context.Context, username, password string) (*model.User, error) {
	id, err := us.GenerateUserId()
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to generate user id").WithError(err)
	}

	password, err = us.EncryptPassword(password)
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to generate user id").WithError(err)
	}

	return &model.User{
		Id:       id,
		Username: username,
		Password: password,
	}, nil
}

func (us *userService) IsBaned(ctx context.Context, uid int64) bool {
	key := us.cache.UserBanedKey(uid)
	return us.cache.IsExist(ctx, key)
}

func (us *userService) CheckPassword(ctx context.Context, ePwd, pwd string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(ePwd), []byte(pwd)); err != nil {
		return errno.NewErrNo(errno.ServiceWrongPassword, "wrong password")
	}
	return nil
}

func (us *userService) UserLogin(ctx context.Context, uid int64) (*model.Token, error) {
	access_token, refresh_token, err := util.CreateAllToken(uid)
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to create token").WithError(err)
	}

	key := us.cache.UserLoginKey(uid)
	err = us.cache.SetUserLogin(ctx, key, refresh_token)
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalRedisErrorCode, "failed to save refresh token to redis").WithError(err)
	}

	return &model.Token{
		AccessToken:       access_token,
		AccessExpireTime:  time.Now().Add(constants.AccessTokenTTL).Unix(),
		RefreshToken:      refresh_token,
		RefreshExpireTime: time.Now().Add(constants.RefreshTokenTTL).Unix(),
	}, nil
}
