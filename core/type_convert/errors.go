package typeconvert

import "errors"

var (
	ErrStringToDateConvertFailed = errors.New("failed to convert string to date")
)

func IsStringToDateConvertFailedError(err error) bool {
	return errors.Is(err, ErrStringToDateConvertFailed)
}
