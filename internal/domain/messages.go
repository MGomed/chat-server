package domain

import (
	"time"

	api "github.com/MGomed/chat_server/pkg/chat_api"
)

// ChatMember domain version of api.ChatMember
type ChatMember struct {
	Name  string
	Email string
}

// ChatInfo domain version of api.ChatInfo
type ChatInfo struct {
	Name    string
	Members []ChatMember
}

// CreateRequest domain version of api.CreateRequest
type CreateRequest struct {
	Info *ChatInfo
}

// CreateReqFromAPIToDomain converts api.CreateRequest to domain.CreateRequest
func CreateReqFromAPIToDomain(req *api.CreateRequest) *CreateRequest {
	members := make([]ChatMember, 0, len(req.Members))
	for _, member := range req.Members {
		members = append(members, ChatMember{
			Name:  member.Name,
			Email: member.Email,
		})
	}

	return &CreateRequest{
		Info: &ChatInfo{
			Name:    req.Name,
			Members: members,
		},
	}
}

// CreateResponse domain version of api.CreateResponse
type CreateResponse struct {
	ID int64
}

// CreateRespToAPIFromDomain converts domain.CreateResponse to api.CreateResponse
func CreateRespToAPIFromDomain(resp *CreateResponse) *api.CreateResponse {
	return &api.CreateResponse{
		Id: resp.ID,
	}
}

// DeleteRequest domain version of api.DeleteRequest
type DeleteRequest struct {
	ID int64
}

// DeleteReqFromAPIToDomain converts api.DeleteRequest to domain.DeleteRequest
func DeleteReqFromAPIToDomain(req *api.DeleteRequest) *DeleteRequest {
	return &DeleteRequest{
		ID: req.Id,
	}
}

// SendRequest domain version of api.SendRequest
type SendRequest struct {
	From      string
	Text      string
	Timestamp time.Time
}

// SendReqFromAPIToDomain converts api.SendRequest to domain.SendRequest
func SendReqFromAPIToDomain(req *api.SendRequest) *SendRequest {
	return &SendRequest{
		From:      req.From,
		Text:      req.Text,
		Timestamp: req.Timestamp.AsTime(),
	}
}
