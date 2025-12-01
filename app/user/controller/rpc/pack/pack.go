package pack

import (
	domainModel "myreel/app/user/domain/model"
	"myreel/kitex_gen/user"
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

func BuildUserWithToken(u *domainModel.User, t *domainModel.Token) *user.UserWithToken {
	return &user.UserWithToken{
		User:  BuildUser(u),
		Token: BuildToken(t),
	}
}
