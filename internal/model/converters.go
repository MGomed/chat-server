package model

import "github.com/MGomed/chat_server/pkg/chat_api"

// ToChatInfoFromAPI converts api.ChatInfo to ChatInfo
func ToChatInfoFromAPI(chat *chat_api.ChatInfo) *ChatInfo {
	return &ChatInfo{
		Name:    chat.Name,
		Members: ToChatMembersFromAPI(chat.Members),
	}
}

// ToChatMembersFromAPI converts api.ChatMember to ChatMember
func ToChatMembersFromAPI(members []*chat_api.ChatMember) []ChatMember {
	res := make([]ChatMember, len(members))
	for _, member := range members {
		res = append(res, ChatMember{
			Name:  member.Name,
			Email: member.Email,
		})
	}

	return res
}

// ToMessageInfoFromAPI converts api.MessageInfo to MessageInfo
func ToMessageInfoFromAPI(message *chat_api.MessageInfo) *MessageInfo {
	return &MessageInfo{
		From:      message.From,
		Text:      message.Text,
		Timestamp: message.Timestamp.AsTime(),
	}
}
