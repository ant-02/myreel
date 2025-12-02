package pack

import (
	domainModel "myreel/app/user/domain/model"
	"myreel/kitex_gen/model"
	"myreel/kitex_gen/user"
	"myreel/pkg/upyun"
)

func BuildUser(u *domainModel.User) *user.User {
	return &user.User{
		Id:        u.Id,
		Username:  u.Username,
		Password:  u.Password,
		AvatarUrl: u.AvatarUrl,
	}
}

func BuildToken(t *domainModel.Token) *user.Token {
	return &user.Token{
		AccessToken:       t.AccessToken,
		AccessExpireTime:  t.AccessExpireTime,
		RefreshToken:      t.RefreshToken,
		RefreshExpireTime: t.RefreshExpireTime,
	}
}

func BuildUpyunToken(t *upyun.UpyunToken) *model.UpyunToken {
	return &model.UpyunToken{
		Policy:        t.Policy,
		Authorization: t.Authorization,
		Bucket:        t.Bucket,
	}
}

func BuildUserWithToken(u *domainModel.User, t *domainModel.Token) *user.UserWithToken {
	return &user.UserWithToken{
		User:  BuildUser(u),
		Token: BuildToken(t),
	}
}
