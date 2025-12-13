package pack

import (
	domainModel "myreel/app/chat/domain/model"
	"myreel/kitex_gen/chat"
)

func BuildMessage(m *domainModel.Message) *chat.Message {
	if m == nil {
		return nil
	}

	return &chat.Message{
		Id:        m.ID,
		SenderId:  m.SenderID,
		TargetId:  m.TargetID,
		Content:   m.Content,
		CreatedAt: m.CreatedAt,
	}
}

func BuildMessages(ms []*domainModel.Message) []*chat.Message {
	l := len(ms)
	if l == 0 {
		return nil
	}

	messages := make([]*chat.Message, l)
	for i, v := range ms {
		messages[i] = BuildMessage(v)
	}
	return messages
}

func BuildPagination(p *domainModel.Pagination) *chat.Pagination {
	return &chat.Pagination{
		NextCursor: p.NextCursor,
		PrevCursor: p.PrevCursor,
		Total:      p.Total,
	}
}

func BuildMessageList(ms []*chat.Message, pagination *chat.Pagination) *chat.MessageList {
	return &chat.MessageList{
		Messages:   ms,
		Pagination: pagination,
	}
}
