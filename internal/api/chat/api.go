package chat

import (
	service "github.com/MGomed/chat_server/internal/service"
	chat_api "github.com/MGomed/chat_server/pkg/chat_api"
)

// API implements UserAPI grpc server
type API struct {
	chat_api.UnimplementedChatAPIServer

	service service.Service
}

// NewAPI is api struct constructor
func NewAPI(service service.Service) *API {
	return &API{
		service: service,
	}
}
