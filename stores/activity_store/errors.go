package activitystore

import "errors"

var (
	NotFoundError = errors.New("Specified id was not found")
)

func IsNotFoundError(err error) bool {
	return errors.Is(err, NotFoundError)
}
