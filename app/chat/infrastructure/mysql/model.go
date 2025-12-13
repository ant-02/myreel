package mysql

import (
	"myreel/pkg/constants"
	"time"

	"gorm.io/gorm"
)

// ChatType 聊天类型
type ChatType int64

const (
	ChatTypePrivate ChatType = 0 // 私聊
	ChatTypeGroup   ChatType = 1 // 群聊
)

// Message 聊天消息
type Message struct {
	ID             int64          `gorm:"type:bigint;primaryKey"`
	ConversationID string         `gorm:"type:varchar(100);not null"`
	SenderID       int64          `gorm:"type:bigint;not null"`
	TargetID       int64          `gorm:"type:bigint;not null"` // 私聊=receiver, 群聊=group_id
	ChatType       ChatType       `gorm:"type:int(2);not null"`
	Content        string         `gorm:"not null;type:text"`
	CreatedAt      time.Time      `gorm:"autoCreateTime"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

// TableName 指定表名
func (Message) TableName() string {
	return constants.MessageTableName
}
