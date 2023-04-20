package inputhandler

import "errors"

var (
	ErrTypeNotFound = errors.New("The type is not in the list of available types")
)

func IsErrTypeNotFoundError(err error) bool {
	return errors.Is(err, ErrTypeNotFound)
}
