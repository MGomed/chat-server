package chat

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	consts "github.com/MGomed/chat_server/consts"
	errors "github.com/MGomed/chat_server/internal/repository/errors"
	db "github.com/MGomed/chat_server/pkg/client/db"
)

// DeleteChat deletes a chat from Postgres DB
func (r *repository) DeleteChat(ctx context.Context, id int64) error {
	builder := sq.Delete(consts.ChatTable).
		Where(sq.Eq{consts.ChatIDColumn: id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%w - %v : %w", errors.ErrQueryBuild, query, err)
	}

	q := db.Query{
		Name:     "chat_server.Delete",
		QueryRaw: query,
	}

	res, err := r.dbc.DB().Exec(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("%w - %v : %w", errors.ErrQueryExecute, query, err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("%w with id: %v", errors.ErrNoSuchChat, id)
	}

	return nil
}
