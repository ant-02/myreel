package pack

import (
	"encoding/json"
	"myreel/app/chat/domain/model"
	"myreel/pkg/errno"
	"strconv"
)

func MessageToMap(v *model.Message) (map[string]interface{}, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, errno.Errorf(errno.InternalRedisErrorCode, "redis: failed to marshal message: %v", err)
	}

	var fields map[string]interface{}
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, errno.Errorf(errno.InternalRedisErrorCode, "redis: failed to unmarshal map: %v", err)
	}

	return fields, nil
}

func MapToMessage(fields map[string]string) (*model.Message, error) {
	if len(fields) == 0 {
		return nil, nil
	}

	parseI64 := func(key string) (int64, error) {
		if s, ok := fields[key]; ok {
			return strconv.ParseInt(s, 10, 64)
		}
		return 0, nil 
	}

	// 解析各字段
	id, err := parseI64("Id")
	if err != nil {
		return nil, err
	}

	senderID, err := parseI64("SenderID")
	if err != nil {
		return nil, err
	}

	targetID, err := parseI64("TargetID")
	if err != nil {
		return nil, err
	}

	chatTypeInt, err := parseI64("ChatType")
	if err != nil {
		return nil, err
	}
	chatType := model.ChatType(chatTypeInt)

	createdAt, err := parseI64("CreatedAt")
	if err != nil {
		return nil, err
	}

	// 构造 Message
	msg := &model.Message{
		ID:             id,
		ConversationID: fields["ConversationID"],
		SenderID:       senderID,
		TargetID:       targetID,
		ChatType:       chatType,
		Content:        fields["Content"],
		CreatedAt:      createdAt,
	}

	return msg, nil
}
