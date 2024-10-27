package model

import "time"

// ChatInfo represent api.ChatInfo object
type ChatInfo struct {
	Name    string
	Members []ChatMember
}

// ChatMember represent api.ChatMember object
type ChatMember struct {
	Name  string
	Email string
}

// MessageInfo represent api.MessageInfo object
type MessageInfo struct {
	From      string
	Text      string
	Timestamp time.Time
}
