package domain

import (
	"time"

	api "github.com/MGomed/chat_server/pkg/chat_api"
)

// CreateRequest domain version of api.CreateRequest
type CreateRequest struct {
	Usernames []string
}

// CreateReqFromAPIToDomain converts api.CreateRequest to domain.CreateRequest
func CreateReqFromAPIToDomain(req *api.CreateRequest) *CreateRequest {
	return &CreateRequest{
		Usernames: req.Usernames,
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

// SendMsgReqFromAPIToDomain converts api.SendRequest to domain.SendRequest
func SendReqFromAPIToDomain(req *api.SendRequest) *SendRequest {
	return &SendRequest{
		From:      req.From,
		Text:      req.Text,
		Timestamp: req.Timestamp.AsTime(),
	}
}
