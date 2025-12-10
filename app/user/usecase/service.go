package usecase

import (
	"context"
	"errors"
	"myreel/app/user/domain/model"
	"myreel/pkg/errno"
	"myreel/pkg/upyun"
)

func (uc *useCase) Register(ctx context.Context, username, password string) error {
	_, err := uc.db.GetUserByUsername(ctx, username)
	if err == nil {
		return errno.NewErrNo(errno.ServiceUserExist, "user has been created")
	}
	if !errors.Is(err, errno.UserNotFound) {
		return err
	}

	u, err := uc.svc.UserRegister(ctx, username, password)
	if err != nil {
		return err
	}

	if err := uc.db.CreateUser(ctx, u); err != nil {
		return err
	}

	return nil
}

func (uc *useCase) Login(ctx context.Context, username, password string) (*model.User, *model.Token, error) {
	u, err := uc.db.GetUserByUsername(ctx, username)
	if err != nil {
		if !errors.Is(err, errno.UserNotFound) {
			return nil, nil, errno.UserNotFound
		}
		return nil, nil, err
	}

	if uc.svc.IsBaned(ctx, u.Id) {
		return nil, nil, errno.UserIsBaned
	}
	if err := uc.svc.CheckPassword(ctx, u.Password, password); err != nil {
		return nil, nil, err
	}

	token, err := uc.svc.UserLogin(ctx, u.Id)
	if err != nil {
		return nil, nil, err
	}

	u.Password = ""

	return u, token, nil
}

func (uc *useCase) GetUserById(ctx context.Context, uid int64) (*model.User, error) {
	return uc.svc.GetUserById(ctx, uid)
}

func (uc *useCase) Refresh(ctx context.Context, token string, uid int64) (string, error) {
	return uc.svc.Refresh(ctx, token, uid)
}

func (uc *useCase) GetLoadToken(ctx context.Context, suffix string, uid int64) (*upyun.UpyunToken, error) {
	return uc.svc.GetUploadToken(ctx, suffix, uid)
}

func (uc *useCase) SetAvatar(ctx context.Context, uid int64, url string) error {
	return uc.svc.SetAvatar(ctx, uid, url)
}

func (uc *useCase) GetUserIdByUsername(ctx context.Context, username string) (int64, error) {
	return uc.db.GetUserIdByUserName(ctx, username)
}

func (uc *useCase) GetUsersByIds(ctx context.Context, ids []int64) ([]*model.UserProfile, error) {
	return uc.svc.GetUsersByIds(ctx, ids)
}
