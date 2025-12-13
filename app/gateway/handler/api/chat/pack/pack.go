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

type SendRequest struct {
	TargetId int64  `json:"targetId"`
	Content  string `json:"content"`
}

type HistoryRequest struct {
	TargetId int64 `json:"targetId"`
	Cursor   int64 `json:"cursor"`
	Limit    int64 `json:"limit"`
}

type UnreadRequest struct {
	TargetId int64 `json:"targetId"`
}

func MarshalWSMessage(respMsg *WSMessage) ([]byte, error) {
	resp, err := json.Marshal(respMsg)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func ParseWSMessage(reqMsg []byte) (*WSMessage, error) {
	var msg WSMessage
	if err := json.Unmarshal(reqMsg, &msg); err != nil {
		return nil, errno.Errorf(errno.InternalServiceErrorCode, "failed to parse request message: %v", err)
	}
	return &msg, nil
}

func ParseSendRequest(msg interface{}) (*SendRequest, error) {
	dataBytes, err := json.Marshal(msg)
	if err != nil {
		return nil, errno.Errorf(errno.InternalServiceErrorCode, "failed to parse send request: %v", err)
	}

	var sendReq SendRequest
	if err := json.Unmarshal(dataBytes, &sendReq); err != nil {
		return nil, errno.Errorf(errno.InternalServiceErrorCode, "failed to unmarshal send request: %v", err)
	}
	return &sendReq, nil
}

func ParseHistoryRequest(msg interface{}) (*HistoryRequest, error) {
	dataBytes, err := json.Marshal(msg)
	if err != nil {
		return nil, errno.Errorf(errno.InternalServiceErrorCode, "failed to parse history request: %v", err)
	}

	var historyReq HistoryRequest
	if err := json.Unmarshal(dataBytes, &historyReq); err != nil {
		return nil, errno.Errorf(errno.InternalServiceErrorCode, "failed to unmarshal history request: %v", err)
	}
	return &historyReq, nil
}

func ParseUnreadRequest(msg interface{}) (*UnreadRequest, error) {
	dataBytes, err := json.Marshal(msg)
	if err != nil {
		return nil, errno.Errorf(errno.InternalServiceErrorCode, "failed to parse unread request: %v", err)
	}

	var unreadReq UnreadRequest
	if err := json.Unmarshal(dataBytes, &unreadReq); err != nil {
		return nil, errno.Errorf(errno.InternalServiceErrorCode, "failed to unmarshal unread request: %v", err)
	}
	return &unreadReq, nil
}
