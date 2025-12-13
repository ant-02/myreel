package mysql

import (
	"myreel/pkg/constants"
	"time"

	"gorm.io/gorm"
)

type Follow struct {
	Id            int64          `gorm:"type:bigint;primaryKey"`
	FolloweringId int64          `gorm:"type:bigint;not null"`
	FolloweredId  int64          `gorm:"type:bigint;not null"`
	Status        int64          `gorm:"type:int(2);noy null;default:0"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type GroupMember struct {
	Id        int64     `gorm:"type:bigint;primaryKey"`
	GroupId   int64     `gorm:"type:bigint;not null"`
	UserId    int64     `gorm:"type:bigint;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type Group struct {
	Id        int64          `gorm:"type:bigint;primaryKey"`
	Name      string         `gorm:"type:varchar(100);not null"`
	CreatorId int64          `gorm:"type:bigint;not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type FolloweredIdWithTime struct {
	FolloweredId int64
	CreatedAt    time.Time
}

func (Follow) TableName() string {
	return constants.FollowTableName
}

func (Group) TableName() string {
	return constants.GroupTableName
}

func (GroupMember) TableName() string {
	return constants.GroupMemberName
}
