package errors

import "errors"

// ErrNameLenInvalid is name validation error
var ErrNameLenInvalid = errors.New("name's length should be between 2 and 32")
