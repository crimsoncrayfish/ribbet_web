package activity

import "errors"

var (
	ErrIdNotFound           = errors.New("Activity with specified id was not found")
	ErrCloseApplication     = errors.New("Safely shutdown application")
	ErrActivityTypeNotFound = errors.New("Activity type does not match expected values")
)

func IsNotFoundError(err error) bool {
	return errors.Is(err, ErrIdNotFound)
}

func IsCloseApplicationError(err error) bool {
	return errors.Is(err, ErrCloseApplication)
}

func IsActivityTypeNotFoundError(err error) bool {
	return errors.Is(err, ErrActivityTypeNotFound)
}
