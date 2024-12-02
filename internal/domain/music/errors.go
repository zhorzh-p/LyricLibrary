package music

import (
	"errors"
	"fmt"
)

var (
	ErrSongAlreadyExists = errors.New("song already exists")
)

type ErrWrongDateFormat struct {
	Date  string
	Cause error
}

func (e ErrWrongDateFormat) Error() string {
	return fmt.Sprintf("failed to parse release date %q. Cause %s", e.Date, e.Cause)
}

func NewErrWrongDateFormat(date string, cause error) error {
	return ErrWrongDateFormat{date, cause}
}
