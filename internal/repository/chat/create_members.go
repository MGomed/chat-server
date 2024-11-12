package chat

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	consts "github.com/MGomed/chat_server/consts"
	service_model "github.com/MGomed/chat_server/internal/model"
	errors "github.com/MGomed/chat_server/internal/repository/errors"
	db "github.com/MGomed/common/pkg/client/db"
)

// CreateMembers creates members in Postgres DB
func (r *repository) CreateMembers(ctx context.Context, chatID int64, members []service_model.ChatMember) error {
	builder := sq.Insert(consts.ChatMemberTable).
		PlaceholderFormat(sq.Dollar).
		Columns(consts.ChatMembersChatIDColumn, consts.ChatMemberNameColumn, consts.ChatMemberEmailColumn)

	for _, member := range members {
		builder = builder.Values(chatID, member.Name, member.Email)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%w - %v : %w", errors.ErrQueryBuild, query, err)
	}

	q := db.Query{
		Name:     "chat_member.Create",
		QueryRaw: query,
	}

	_, err = r.dbc.DB().Query(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("%w - %v : %w", errors.ErrQueryExecute, query, err)
	}

	return nil
}
