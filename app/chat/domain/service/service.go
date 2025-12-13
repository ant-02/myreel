package service

import (
	"context"
	"myreel/app/chat/domain/model"
	"myreel/pkg/constants"
	"myreel/pkg/errno"
	"strconv"
)

func (cs *chatService) GenerateConversationID(chatType model.ChatType, id1, id2 int64) string {
	if chatType == model.ChatTypePrivate {
		// 私聊：用 min+max 保证唯一且无序
		if id1 > id2 {
			id1, id2 = id2, id1
		}
		return "p_" + strconv.FormatInt(id1, 10) + "_" + strconv.FormatInt(id2, 10)
	}
	// 群聊：直接用 group_id
	return "g_" + strconv.FormatInt(id1, 10)
}

func (cs *chatService) GenerateMessageId() (int64, error) {
	id, err := cs.sf.Generate()
	if err != nil {
		return 0, err
	}
	return id, err
}

func (cs *chatService) CreateMessage(ctx context.Context, msg *model.Message) error {
	hk, uk, mk := cs.cache.HistoryKey(msg.ConversationID), cs.cache.UnreadKey(msg.ConversationID), cs.cache.MessageKey(msg.ID)

	err := cs.cache.AddMessageId(ctx, hk, float64(msg.CreatedAt), msg.ID)
	if err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to add message id to history zset").WithError(err)
	}

	err = cs.cache.AddMessageId(ctx, uk, float64(msg.CreatedAt), msg.ID)
	if err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to add message id to unread zset").WithError(err)
	}

	err = cs.cache.AddMessageWithTTL(ctx, mk, msg, constants.RedisChatMessageExpireTime)
	if err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to add message to hash").WithError(err)
	}

	err = cs.db.CreateMessage(ctx, msg)
	if err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to create message").WithError(err)
	}
	return nil
}
