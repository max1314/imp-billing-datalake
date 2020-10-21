// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package errors

import (
	"strings"
	"sync"

	"google.golang.org/grpc/codes"
)

// ComboError is a helper struct for returning an error that combines multiple other errors.
type ComboError struct {
	mu       sync.Mutex
	errors   []error
	lastCode codes.Code
}

// NewCombo returns an empty ComboError.
func NewCombo() *ComboError {
	return &ComboError{}
}

// Add appends the error and sets the Code to that of the error. This method is thread safe.
//
// For example
//   combo := NewCombo()
//   combo.Add(errors.New(nil, errors.Internal, "internal error"))
//   combo.Add(errors.New(nil, errors.NotFound, "not found"))
//   // combo.Code() returns errors.NotFound
func (c *ComboError) Add(err error) {
	if err == nil {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	if e, ok := err.(frame); ok {
		c.lastCode = e.code
		c.errors = append(c.errors, e)
		return
	}
	if e, ok := err.(list); ok {
		c.lastCode = e.code
		c.errors = append(c.errors, e)
		return
	}
	e := newFrame(nil, Code(err), Desc(err))
	c.lastCode = e.code
	c.errors = append(c.errors, e)
}

// AddErrors appends all errors from another ComboError and sets the Code to the other combo's. Method is thread safe.
//
// For example
//   combo := NewCombo()
//   combo2 := NewCombo()
//   combo2.Add(errors.New(nil, errors.PermissionDenied, "permission denied"))
//   combo.AddErrors(combo2)
//   // combo.Code() returns errors.PermissionDenied
//   combo.Add(NewCombo())
//   // combo.Code() still returns errors.PermissionDenied
func (c *ComboError) AddErrors(other *ComboError) {
	if other == nil {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	c.errors = append(c.errors, other.errors...)
	if other.Code() != OK {
		c.lastCode = other.Code()
	}
}

// messages returns the concatenated error messages
func (c *ComboError) messages() string {
	outs := []string{}
	for _, err := range c.errors {
		outs = append(outs, Desc(err))
	}
	return strings.Join(outs, "\n\t")
}

// details include the file name and line number of the errors
func (c *ComboError) details() string {
	outs := []string{}
	for _, err := range c.errors {
		outs = append(outs, Details(err))
	}
	return strings.Join(outs, "\n\t")
}

// Code implements the `ErrorCoder` interface
func (c *ComboError) Code() codes.Code {
	return c.lastCode
}

// Len returns the total number of error that is contained within this combo. This method is thread safe.
func (c *ComboError) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.errors)
}

// NilOrError returns a combined Error
func (c *ComboError) NilOrError() error {
	if len(c.errors) == 0 {
		return nil
	}
	return list{
		frame: newFrame(nil, c.lastCode, c.messages()),
		Errs:  c.errors,
	}
}
