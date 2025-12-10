package model

type Follow struct {
	Id            int64
	FolloweringId int64
	FolloweredId  int64
	Status        int64
	CreatedAt     int64
}
