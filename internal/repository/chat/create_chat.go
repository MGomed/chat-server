package chat

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	db "github.com/MGomed/chat_server/pkg/client/db"
)

func (r *repository) CreateChat(ctx context.Context, name string) (int64, error) {
	builder := sq.Insert(chatTable).
		PlaceholderFormat(sq.Dollar).
		Columns(chatNameColumn).
		Values(name).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errQueryBuild, query, err)
	}

	q := db.Query{
		Name:     "chat_server.Create",
		QueryRaw: query,
	}

	var chatID int64
	err = r.dbc.DB().QueryRow(ctx, q, args...).Scan(&chatID)
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errQueryExecute, query, err)
	}

	return chatID, nil
}
