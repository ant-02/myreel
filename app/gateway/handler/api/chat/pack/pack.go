package pack

import (
	"encoding/json"
	"myreel/pkg/errno"
)

const (
	TypePrivateMessage int = iota // 私聊发送消息
	TypePrivateHistory            // 获取私聊历史记录
	TypePrivateUnread             // 获取私聊未读消息
	TypeGroupMessage              // 群聊发送消息
	TypeGroupHistory              // 获取群聊历史记录
)

type WSMessage struct {
	Type int         `json:"type"`
	Data interface{} `json:"data"`
}

type PrivateMsg struct {
	TargetId int64 `json:"targetId"`
	Content  string `json:"content"`
}

func ParseMessage(reqMsg []byte) (*WSMessage, error) {
	var msg WSMessage
	if err := json.Unmarshal(reqMsg, &msg); err != nil {
		return nil, errno.Errorf(errno.InternalServiceErrorCode, "failed to parse request message: %v", err)
	}
	return &msg, nil
}

func ParsePrivateMsg(msg interface{}) (*PrivateMsg, error) {
	dataBytes, err := json.Marshal(msg)
	if err != nil {
		return nil, errno.Errorf(errno.InternalServiceErrorCode, "failed to parse private message: %v", err)
	}

	var privateMsg PrivateMsg
	if err := json.Unmarshal(dataBytes, &privateMsg); err != nil {
		return nil, errno.Errorf(errno.InternalServiceErrorCode, "failed to unmarshal private message: %v", err)
	}
	return &privateMsg, nil
}
