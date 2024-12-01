package errors

import "errors"

// Access errors
var (
	ErrMetadataNotProvided = errors.New("metadata is not provided")
	ErrHeaderNotProvided   = errors.New("authorization header is not provided")
	ErrHeaderWrongFormat   = errors.New("invalid authorization header format")
)
