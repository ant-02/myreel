package model

// ChatType 聊天类型
type ChatType int64

const (
	ChatTypePrivate ChatType = 0 // 私聊
	ChatTypeGroup   ChatType = 1 // 群聊
)

type Message struct {
	ID             int64
	ConversationID string
	SenderID       int64
	TargetID       int64
	ChatType       ChatType
	Content        string
	CreatedAt      int64
}
