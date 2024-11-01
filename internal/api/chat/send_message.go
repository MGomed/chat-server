package chat

import (
	"context"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	service_conv "github.com/MGomed/chat_server/internal/converters"
	chat_api "github.com/MGomed/chat_server/pkg/chat_api"
)

// SendMessage send user's message to chat
func (a *API) SendMessage(ctx context.Context, req *chat_api.SendRequest) (*emptypb.Empty, error) {
	err := a.service.SendMessage(ctx, service_conv.ToMessageInfoFromAPI(req.Info))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
