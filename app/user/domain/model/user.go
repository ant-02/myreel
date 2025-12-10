package model

type User struct {
	Id        int64
	Username  string
	Password  string
	AvatarUrl string
}

type Token struct {
	AccessToken       string
	AccessExpireTime  int64
	RefreshToken      string
	RefreshExpireTime int64
}

type UserProfile struct {
	Id        int64
	Username  string
	AvatarUrl string
}
