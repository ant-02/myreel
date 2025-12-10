package service

import (
	"context"
	"fmt"
	"myreel/app/user/domain/model"
	"myreel/config"
	"myreel/pkg/constants"
	"myreel/pkg/errno"
	"myreel/pkg/upyun"
	"myreel/pkg/util"
	"strconv"
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

func (us *userService) GetUserById(ctx context.Context, uid int64) (*model.User, error) {
	u, err := us.db.GetUserById(ctx, uid)
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get user by id").WithError(err)
	}
	return u, nil
}

func (us *userService) GetUploadToken(ctx context.Context, suffix string, uid int64) (*upyun.UpyunToken, error) {
	saveKey := fmt.Sprintf("%s/%s/%d%s", constants.UpyunUserAvaterPath, time.Now().Format("2006/01/02"), uid, suffix)
	up, err := upyun.GeneratePolicyAndSignature(uid, saveKey, fmt.Sprintf("%s%s", config.Upyun.NotifyUrl, constants.UpyunUserAvatarNotifyPath), strconv.FormatInt(uid, 10))
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get upyun token").WithError(err)
	}

	return up, nil
}

func (us *userService) Refresh(ctx context.Context, token string, uid int64) (string, error) {
	key := us.cache.UserLoginKey(uid)

	if exist := us.cache.IsExist(ctx, key); !exist {
		return "", errno.NewErrNo(errno.InternalServiceErrorCode, "reject to issue access token")
	}

	refreshToken, err := us.cache.GetUserLogin(ctx, key)
	if err != nil {
		return "", errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get refresh token").WithError(err)
	}

	if refreshToken != token {
		return "", errno.NewErrNo(errno.InternalServiceErrorCode, "reject to issue access token")
	}

	accessToken, err := util.CreateToken(constants.TypeAccessToken, uid)
	if err != nil {
		return "", errno.NewErrNo(errno.InternalServiceErrorCode, "failed to create token").WithError(err)
	}

	return accessToken, nil
}

func (us *userService) SetAvatar(ctx context.Context, uid int64, url string) error {
	if err := us.db.SetAvatar(ctx, uid, fmt.Sprintf("%s%s", config.Upyun.Domain, url)); err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to set user avatar").WithError(err)
	}
	return nil
}

func (us *userService) GetUserIdByUsername(ctx context.Context, username string) (int64, error) {
	id, err := us.db.GetUserIdByUserName(ctx, username)
	if err != nil {
		return 0, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get user id by username").WithError(err)
	}
	return id, nil
}

func (us *userService) GetUsersByIds(ctx context.Context, ids []int64) ([]*model.UserProfile, error) {
	users, err := us.db.GetUsersByIds(ctx, ids)
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "faile to get users by ids").WithError(err)
	}
	return users, nil
}
