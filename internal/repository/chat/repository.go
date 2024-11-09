package chat

import (
	db "github.com/MGomed/common/pkg/client/db"
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
