package errors

import "fmt"

type dataLakeError struct {
	err error
	msg string
}

func Wrap(err error, msg string) *dataLakeError{
	return &dataLakeError{err: err, msg: msg}
}

func New(msg string) *dataLakeError{
	return &dataLakeError{err: nil, msg: msg }
}

func (e *dataLakeError) Error() string {
	return fmt.Sprintf("%s: %s", e.err, e.msg)
}
