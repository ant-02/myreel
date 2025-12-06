package mysql

import (
	"myreel/pkg/constants"
	"time"
)

type Comment struct {
	Id         int64     `gorm:"type:bigint;primaryKey"`
	VideoId    int64     `gorm:"type:bigint"`
	Uid        int64     `gorm:"type:bigint"`
	ParentId   int64     `gorm:"type:bigint;default:0"`
	LikeCount  int64     `gorm:"type:int;default:0"`
	ChildCount int64     `gorm:"type:int;default:0"`
	Content    string    `gorm:"type:varchar(1000);null not"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	DeletedAt  time.Time `gorm:"type:datetime;default:null"`
}

func (Comment) TableName() string {
	return constants.CommentTableName
}
