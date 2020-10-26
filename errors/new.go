package errors

import "fmt"

type DataLakeError struct {
	Err error
	msg string
}

func New(err error, msg string) *DataLakeError {
	return &DataLakeError{Err: err, msg: msg}
}

func (e *DataLakeError) Error() string {
	return fmt.Sprintf("%s: %s", e.Err, e.msg)
}
