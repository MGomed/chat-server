package chat

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	db "github.com/MGomed/chat_server/pkg/client/db"
)

func (r *repository) DeleteChat(ctx context.Context, id int64) error {
	builder := sq.Delete(chatTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{chatIDColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%w - %v : %w", errQueryBuild, query, err)
	}

	q := db.Query{
		Name:     "chat_server.Delete",
		QueryRaw: query,
	}

	res, err := r.dbc.DB().Exec(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("%w - %v : %w", errQueryExecute, query, err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("%w with id: %v", errNoSuchChat, id)
	}

	return nil
}
