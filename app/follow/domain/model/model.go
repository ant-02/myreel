package model

import "time"

type Follow struct {
	Id            int64
	FolloweringId int64
	FolloweredId  int64
	Status        int64
	CreatedAt     int64
}

type FolloweredIdWithTime struct {
	FolloweredId int64
	CreatedAt    time.Time
}

type FolloweringIdWithTime struct {
	FolloweringId int64
	CreatedAt     time.Time
}

type UserProfile struct {
	Id        int64
	Username  string
	AvatarUrl string
}

type Pagination struct {
	NextCursor int64
	PrevCursor int64
	Total      int64
}

type GroupMember struct {
	Id        int64
	GroupId   int64
	UserId    int64
	CreatedAt int64
}

type Group struct {
	Id        int64
	Name      string
	CreatorId int64
	CreatedAt int64
}
