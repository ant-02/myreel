package mysql

import (
	"context"
	"myreel/app/chat/domain/model"
	"myreel/app/chat/domain/repository"
	"myreel/pkg/errno"

	"gorm.io/gorm"
)

type chatDB struct {
	client *gorm.DB
}

func NewChatDB(client *gorm.DB) repository.ChatDB {
	return &chatDB{client: client}
}

func (db *chatDB) Magrate() error {
	if err := db.client.AutoMigrate(&Message{}); err != nil {
		return errno.NewErrNo(errno.InternalDatabaseErrorCode, "mysql: failed to auto magrate chat model")
	}
	return nil
}

func (db *chatDB) CreateMessage(ctx context.Context, msg *model.Message) error {
	m := &Message{
		ID:             msg.ID,
		ConversationID: msg.ConversationID,
		SenderID:       msg.SenderID,
		TargetID:       msg.TargetID,
		ChatType:       ChatType(msg.ChatType),
		Content:        msg.Content,
	}
	err := db.client.WithContext(ctx).Create(&m).Error
	if err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to create message: %v", err)
	}
	return nil
}
