package consts

import "time"

const ContextTimeout = 15 * time.Second

const (
	ServerHostEnv = "SERVER_HOST"
	ServerPortEnv = "SERVER_PORT"
)

const (
	DBHostEnv     = "DB_HOST"
	DBPortEnv     = "DB_PORT"
	DBNameEnv     = "POSTGRES_DB"
	DBUserEnv     = "POSTGRES_USER"
	DBPasswordEnv = "POSTGRES_PASSWORD" //nolint: gosec
)

const (
	ChatTable      = "chat"
	ChatIDColumn   = "id"
	ChatNameColumn = "name"

	ChatMemberTable         = "chat_member"
	ChatMembersChatIDColumn = "chat_id"
	ChatMemberNameColumn    = "name"
	ChatMemberEmailColumn   = "email"
)
