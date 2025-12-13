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
	return "g_" + strconv.FormatInt(id2, 10)
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

func (cs *chatService) GetHistoryMessages(ctx context.Context, cursor, limit int64, conversationID string) ([]*model.Message, *model.Pagination, error) {
	hk, uk := cs.cache.HistoryKey(conversationID), cs.cache.UnreadKey(conversationID)
	ids, err := cs.cache.GetMessageIds(ctx, hk, cursor, limit)
	if err != nil {
		return nil, nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get message id list").WithError(err)
	}
	total, err := cs.cache.GetMessageCount(ctx, hk)
	if err != nil {
		return nil, nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get message count").WithError(err)
	}

	l := len(ids)
	msgs := make([]*model.Message, l)
	nextCursor := cursor
	if l > 0 {
		err = cs.cache.RemUnreadMessage(ctx, uk, ids...)
		if err != nil {
			return nil, nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to remove message").WithError(err)
		}
		for i, v := range ids {
			mk := cs.cache.MessageKey(v)
			var msg *model.Message
			var err error
			if exist := cs.cache.IsExist(ctx, mk); exist {
				msg, err = cs.cache.GetMessage(ctx, mk)
				if err != nil {
					return nil, nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get history message").WithError(err)
				}
			} else {
				msg, err = cs.db.GetMessage(ctx, v)
				if err != nil {
					return nil, nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get history message").WithError(err)
				}
				err = cs.cache.AddMessageWithTTL(ctx, mk, msg, constants.RedisChatMessageExpireTime)
				if err != nil {
					return nil, nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to add message to hash").WithError(err)
				}
			}
			msgs[i] = msg
		}
		nextCursor = msgs[l-1].CreatedAt
	}

	return msgs, &model.Pagination{
		NextCursor: nextCursor,
		PrevCursor: cursor,
		Total:      total,
	}, nil
}

func (cs *chatService) GetUnreadMessages(ctx context.Context, conversationID string) ([]*model.Message, *model.Pagination, error) {
	uk := cs.cache.UnreadKey(conversationID)

	total, err := cs.cache.GetMessageCount(ctx, uk)
	if err != nil {
		return nil, nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get message count").WithError(err)
	}
	if total == 0 {
		return nil, &model.Pagination{
			NextCursor: 0,
			PrevCursor: 0,
			Total:      total,
		}, nil
	}
	ids, err := cs.cache.GetMessageIds(ctx, uk, 0, total)
	if err != nil {
		return nil, nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get message id list").WithError(err)
	}

	err = cs.cache.RemUnreadMessage(ctx, uk, ids...)
	if err != nil {
		return nil, nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to remove message").WithError(err)
	}

	l := len(ids)
	msgs := make([]*model.Message, l)
	for i, v := range ids {
		mk := cs.cache.MessageKey(v)
		var msg *model.Message
		var err error
		if exist := cs.cache.IsExist(ctx, mk); exist {
			msg, err = cs.cache.GetMessage(ctx, mk)
			if err != nil {
				return nil, nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get history message").WithError(err)
			}
		} else {
			msg, err = cs.db.GetMessage(ctx, v)
			if err != nil {
				return nil, nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get history message").WithError(err)
			}
			err = cs.cache.AddMessageWithTTL(ctx, mk, msg, constants.RedisChatMessageExpireTime)
			if err != nil {
				return nil, nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to add message to hash").WithError(err)
			}
		}
		msgs[i] = msg
	}

	return msgs, &model.Pagination{
		NextCursor: 0,
		PrevCursor: 0,
		Total:      total,
	}, nil
}
