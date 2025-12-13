package usecase

import (
	"context"
	"myreel/app/chat/domain/model"
	"time"
)

func (uc *useCase) SendMessage(ctx context.Context, senderID, targetID int64, chatType model.ChatType, content string) error {
	id, err := uc.svc.GenerateMessageId()
	if err != nil {
		return err
	}

	conversationID := uc.svc.GenerateConversationID(chatType, senderID, targetID)
	createdAt := time.Now().Unix()

	msg := &model.Message{
		ID:             id,
		ConversationID: conversationID,
		SenderID:       senderID,
		TargetID:       targetID,
		ChatType:       chatType,
		Content:        content,
		CreatedAt:      createdAt,
	}
	return uc.svc.CreateMessage(ctx, msg)
}
