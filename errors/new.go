package errors

import "fmt"

type DataLakeError struct {
	err error
	msg string
}

func Wrap(err error, msg string) *DataLakeError {
	return &DataLakeError{err: err, msg: msg}
}

func New(msg string) *DataLakeError {
	return &DataLakeError{err: nil, msg: msg }
}

func (e *DataLakeError) Error() string {
	return fmt.Sprintf("%s: %s", e.err, e.msg)
}
