package errors

import "errors"

var (
	ErrNameLenInvalid = errors.New("name's length should be between 2 and 32")
)
