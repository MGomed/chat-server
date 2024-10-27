package chat

import (
	"errors"

	db "github.com/MGomed/chat_server/pkg/client/db"
)

const (
	chatTable      = "chat"
	chatIDColumn   = "id"
	chatNameColumn = "name"

	chatMemberTable         = "chat_member"
	chatMembersChatIDColumn = "chat_id"
	chatMemberNameColumn    = "name"
	chatMemberEmailColumn   = "email"
)

var (
	errQueryBuild   = errors.New("failed to build query")
	errQueryExecute = errors.New("failed to execute query")
	errNoSuchChat   = errors.New("chat not found")
)

type repository struct {
	dbc db.Client
}

// NewRepository is adapter struct constructor
func NewRepository(dbc db.Client) *repository {
	return &repository{
		dbc: dbc,
	}
}
