package cache

import (
	"fmt"
	"myreel/pkg/constants"
)

func (cc *chatCache) HistoryKey(conversationID string) string {
	return constants.RedisChatHistoryKey + conversationID
}

func (cc *chatCache) UnreadKey(conversationID string) string {
	return constants.RedisChatUnreadKey + conversationID
}

func (cc *chatCache) MessageKey(id int64) string {
	return fmt.Sprintf("%s%d", constants.RedisChatMessage, id)
}
