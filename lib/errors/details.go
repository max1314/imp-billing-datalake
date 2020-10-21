// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package errors

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

// Details returns information about the stack of
// underlying errors wrapped by err, in the format:
//
// {filename:99: code=x desc=error one}
// {otherfile:55: code=y desc=cause of error one}
//
// The details are found by type-asserting the error to our error type
// Details of the underlying stack are found by
// recursively calling Underlying until it is nil
func Details(err error) string {
	var b bytes.Buffer
	writeDetails(&b, err, 0)
	if b.Len() == 0 {
		return ""
	}
	return b.String()
}

func writeDetails(b *bytes.Buffer, err error, indent int) {
	for err != nil {
		b.WriteString(strings.Repeat("\t", indent))
		b.WriteRune('{')

		if e, ok := err.(sourcer); ok {
			f, l := e.FileAndLine()
			if f != "" {
				fmt.Fprintf(b, "%s:%d: ", f, l)
			}
		}

		fmt.Fprintf(b, "code = %s", Code(err).String())

		if errs, ok := causeList(err); ok {
			b.WriteString(" details = {")
			for _, subErr := range errs {
				b.WriteRune('\n')
				writeDetails(b, subErr, indent+1)
			}
			b.WriteRune('}')
		} else {
			fmt.Fprintf(b, " desc = %s", Desc(err))
		}

		b.WriteRune('}')

		if errFrame, ok := err.(frame); ok && errFrame.stackTrace != "" {
			b.WriteRune('\n')
			fmt.Fprint(b, errFrame.stackTrace)
		}

		err = Cause(err)

		if err != nil {
			b.WriteRune('\n')
		}
	}
}

// Cause returns the next error in the error stack if supported by the underlying error type.
func Cause(err error) error {
	if err == nil {
		return nil
	}
	if ewo, ok := err.(errWithOrigin); ok {
		orig := ewo.Origin()
		if orig == err {
			return nil
		}
		return orig
	}
	if ewc, ok := err.(errWithCause); ok {
		return ewc.Cause()
	}
	v := reflect.ValueOf(err)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}
	v = v.FieldByName("Err")
	if !v.IsValid() {
		return nil
	}
	err, _ = v.Interface().(error)
	return err
}

func causeList(err error) ([]error, bool) {
	if err == nil {
		return nil, false
	}
	v := reflect.ValueOf(err)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, false
	}
	if l := v.FieldByName("Errors"); l.IsValid() {
		v = l
	} else if l := v.FieldByName("Errs"); l.IsValid() {
		v = l
	} else {
		return nil, false
	}
	if v.Kind() != reflect.Slice {
		return nil, false
	}
	errs, ok := v.Interface().([]error)
	return errs, ok
}

type sourcer interface {
	error
	FileAndLine() (string, int)
}

// as found on grpc/transport.ConnectionError
type errWithOrigin interface {
	error
	// returns the original error, which may be this error or an underlying error
	Origin() error
}

// as found in github.com/pkg/errors and github.com/gorilla/securecookie
type errWithCause interface {
	error
	// if implemented, returns the underlying error
	Cause() error
}
