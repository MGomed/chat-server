package chat

import (
	"context"

	service_model "github.com/MGomed/chat_server/internal/model"
)

// Create creates new chat
func (s *service) Create(ctx context.Context, chat *service_model.ChatInfo) (int64, error) {
	id, err := s.repo.CreateChat(ctx, chat.Name)
	if err != nil {
		s.logger.Error("Failed to add chat %v in database: %v", chat.Name, err)

		return 0, err
	}

	if len(chat.Members) > 0 {
		if err := s.repo.CreateMembers(ctx, id, chat.Members); err != nil {
			s.logger.Error("Failed to add chat_members %v in database: %v", chat.Members, err)

			return id, err
		}
	}

	s.logger.Debug("Successfully added chat with id: %v", id)

	return id, nil
}
