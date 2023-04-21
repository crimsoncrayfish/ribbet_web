package inputhandler

import "errors"

var (
	ErrTypeNotFound = errors.New("the type is not in the list of available types")
)

func IsTypeNotFoundError(err error) bool {
	return errors.Is(err, ErrTypeNotFound)
}
