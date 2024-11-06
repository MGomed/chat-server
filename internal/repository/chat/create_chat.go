package chat

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	consts "github.com/MGomed/chat_server/consts"
	errors "github.com/MGomed/chat_server/internal/repository/errors"
	db "github.com/MGomed/chat_server/pkg/client/db"
)

// CreateChat creates a chat in Postgres DB
func (r *repository) CreateChat(ctx context.Context, name string) (int64, error) {
	builder := sq.Insert(consts.ChatTable).
		Columns(consts.ChatNameColumn).
		Values(name).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errors.ErrQueryBuild, query, err)
	}

	q := db.Query{
		Name:     "chat_server.Create",
		QueryRaw: query,
	}

	var chatID int64
	err = r.dbc.DB().QueryRow(ctx, q, args...).Scan(&chatID)
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errors.ErrQueryExecute, query, err)
	}

	return chatID, nil
}
