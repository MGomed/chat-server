package chat

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	service_model "github.com/MGomed/chat_server/internal/model"
	db "github.com/MGomed/chat_server/pkg/client/db"
)

func (r *repository) CreateMembers(ctx context.Context, chatID int64, members ...service_model.ChatMember) error {
	builder := sq.Insert(chatMemberTable).
		PlaceholderFormat(sq.Dollar).
		Columns(chatMembersChatIDColumn, chatMemberNameColumn, chatMemberEmailColumn)

	for _, member := range members {
		builder = builder.Values(chatID, member.Name, member.Email)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%w - %v : %w", errQueryBuild, query, err)
	}

	q := db.Query{
		Name:     "chat_member.Create",
		QueryRaw: query,
	}

	_, err = r.dbc.DB().Query(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("%w - %v : %w", errQueryExecute, query, err)
	}

	return nil
}
