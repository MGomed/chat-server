package chat

import (
	"context"

	service_model "github.com/MGomed/chat_server/internal/model"
)

// Create creates new chat
func (s *service) Create(ctx context.Context, chat *service_model.ChatInfo) (int64, error) {
	var id int64

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		id, errTx = s.repo.CreateChat(ctx, chat.Name)
		if errTx != nil {
			s.logger.Printf("Failed to add chat %v in database: %v", chat.Name, errTx)

			return errTx
		}

		if len(chat.Members) > 0 {
			errTx = s.repo.CreateMembers(ctx, id, chat.Members...)
			if errTx != nil {
				s.logger.Printf("Failed to add chat_members %v in database: %v", chat.Members, errTx)

				return errTx
			}
		}

		s.logger.Printf("Successfully added chat with id: %v", id)

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
