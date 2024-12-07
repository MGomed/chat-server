package consts

import "time"

// ServiceName is a application name
const ServiceName = "chat_service"

// ContextTimeout is timeout for db connecting and server start
const ContextTimeout = 15 * time.Second

// Server env's names
const (
	ServerHostEnv = "SERVER_HOST"
	ServerPortEnv = "SERVER_PORT"
)

// Access Service env's names
const (
	AccessServiceHostEnv = "AUTH_SERVICE_HOST"
	AccessServicePortEnv = "AUTH_SERVICE_PORT"
)

// DB env's names
const (
	DBHostEnv     = "DB_HOST"
	DBPortEnv     = "DB_PORT"
	DBNameEnv     = "POSTGRES_DB"
	DBUserEnv     = "POSTGRES_USER"
	DBPasswordEnv = "POSTGRES_PASSWORD" //nolint: gosec
)

// DB table and columns names
const (
	ChatTable      = "chat"
	ChatIDColumn   = "id"
	ChatNameColumn = "name"

	ChatMemberTable         = "chat_member"
	ChatMembersChatIDColumn = "chat_id"
	ChatMemberNameColumn    = "name"
	ChatMemberEmailColumn   = "email"
)

// AccessPrefix defines access prefix in grpc context
var AccessPrefix = "Bearer "
