package errors

import "fmt"

type DataLakeError struct {
	err error
	msg string
}

func New(err error, msg string) *DataLakeError {
	return &DataLakeError{err: err, msg: msg}
}

func (e *DataLakeError) InnerErr() error {
	return e.err
}

func (e *DataLakeError) Error() string {
	return fmt.Sprintf("%v: %s", e.err, e.msg)
}
