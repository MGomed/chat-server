package chat

import (
	"context"

	service_model "github.com/MGomed/chat_server/internal/model"
)

// SendMessage send user's message to chat
func (s *service) SendMessage(_ context.Context, _ *service_model.MessageInfo) error {
	// TODO some business logic

	return nil
}
