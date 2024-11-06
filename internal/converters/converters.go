package converters

import (
	service_model "github.com/MGomed/chat_server/internal/model"
	chat_api "github.com/MGomed/chat_server/pkg/chat_api"
)

// ToChatInfoFromAPI converts api.ChatInfo to ChatInfo
func ToChatInfoFromAPI(chat *chat_api.ChatInfo) *service_model.ChatInfo {
	if chat == nil {
		return nil
	}

	return &service_model.ChatInfo{
		Name:    chat.Name,
		Members: ToChatMembersFromAPI(chat.Members),
	}
}

// ToChatMembersFromAPI converts api.ChatMember to ChatMember
func ToChatMembersFromAPI(members []*chat_api.ChatMember) []service_model.ChatMember {
	res := make([]service_model.ChatMember, 0, len(members))
	for _, member := range members {
		res = append(res, service_model.ChatMember{
			Name:  member.Name,
			Email: member.Email,
		})
	}

	return res
}

// ToMessageInfoFromAPI converts api.MessageInfo to MessageInfo
func ToMessageInfoFromAPI(message *chat_api.MessageInfo) *service_model.MessageInfo {
	if message == nil {
		return nil
	}

	return &service_model.MessageInfo{
		From:      message.From,
		Text:      message.Text,
		Timestamp: message.Timestamp.AsTime(),
	}
}
