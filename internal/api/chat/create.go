package chat

import (
	"context"

	service_model "github.com/MGomed/chat_server/internal/model"
	chat_api "github.com/MGomed/chat_server/pkg/chat_api"
)

// Create creates new chat
func (a *API) Create(ctx context.Context, req *chat_api.CreateRequest) (*chat_api.CreateResponse, error) {
	id, err := a.service.Create(ctx, service_model.ToChatInfoFromAPI(req.Chat))
	if err != nil {
		return nil, err
	}

	return &chat_api.CreateResponse{
		Id: id,
	}, nil
}
