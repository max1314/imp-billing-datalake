// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package errors

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"improbable.io/lib/sharedflags"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
)

func TestComboError_AddErrors(t *testing.T) {
	for _, spec := range []struct {
		source        *ComboError
		argument      *ComboError
		expectedCount int
		expectedCode  codes.Code
	}{
		{
			source:        NewCombo(),
			argument:      nil,
			expectedCount: 0,
			expectedCode:  codes.OK,
		},
		{
			source:        NewCombo(),
			argument:      NewCombo(),
			expectedCount: 0,
			expectedCode:  codes.OK,
		},
		{
			source:        NewCombo(),
			argument:      &ComboError{errors: []error{New(nil, codes.PermissionDenied, "secondError")}, lastCode: codes.PermissionDenied},
			expectedCount: 1,
			expectedCode:  codes.PermissionDenied,
		},
		{
			source:        &ComboError{errors: []error{New(nil, codes.PermissionDenied, "firstError")}, lastCode: codes.PermissionDenied},
			argument:      NewCombo(),
			expectedCount: 1,
			expectedCode:  codes.PermissionDenied,
		},
		{
			source:        &ComboError{errors: []error{New(nil, codes.Unknown, "firstError")}, lastCode: codes.Unknown},
			argument:      &ComboError{errors: []error{New(nil, codes.PermissionDenied, "secondError")}, lastCode: codes.PermissionDenied},
			expectedCount: 2,
			expectedCode:  codes.PermissionDenied,
		},
		{
			source:        &ComboError{errors: []error{New(nil, codes.Unknown, "firstError")}, lastCode: codes.PermissionDenied},
			argument:      &ComboError{errors: []error{New(nil, codes.PermissionDenied, "secondError")}, lastCode: codes.Unknown},
			expectedCount: 2,
			expectedCode:  codes.Unknown,
		},
	} {
		spec.source.AddErrors(spec.argument)
		assert.Equal(t, spec.expectedCode, spec.source.Code())
		assert.Equal(t, spec.expectedCode, Code(spec.source.NilOrError()))
		assert.Equal(t, spec.expectedCount, spec.source.Len())
	}
}

func TestAddEmptyError(t *testing.T) {
	e := NewCombo()
	e.Add(nil)
	err := e.NilOrError()
	if err != nil {
		println(err.Error())
	}
}

func splitDetailsForTest(t *testing.T, err error) (indents []string, tails []string) {
	details := strings.Split(Details(err), "\n")
	for _, d := range details {
		indents = append(indents, d[:strings.Index(d, "{")])
		tails = append(tails, d[strings.LastIndex(d, "src/improbable.io/")+len("src/improbable.io/"):])
	}
	return indents, tails
}

// Detail(err) will print out
// ```
// {$GOPATH/src/improbable.io/lib/errors/combo_error.go:95: code = NotFound details = {
//   {$GOPATH/src/improbable.io/lib/errors/combo_error_test.go:92: code = Unknown desc = Error}
//   {$GOPATH/src/improbable.io/lib/errors/combo_error_test.go:93: code = NotFound desc = Err1}
//   {$GOPATH/src/improbable.io/lib/errors/combo_error_test.go:94: code = NotFound desc = Err2}}}
// ```

func TestAddExternalErr(t *testing.T) {
	sharedflags.Set.Set("errors_include_debug_info", "true")
	err := NewCombo()
	_, _, line, _ := runtime.Caller(0)  // line+0
	err.Add(fmt.Errorf("Error"))        // line+1
	err.Add(New(nil, NotFound, "Err1")) // line+2
	err.Add(New(nil, NotFound, "Err2")) // line+3
	err2 := err.NilOrError()            // line+4
	indents, tails := splitDetailsForTest(t, err2)
	assert.Equal(t, "", indents[0])
	assert.Equal(t, "\t", indents[1])
	assert.Equal(t, "\t", indents[2])
	assert.Equal(t, "\t", indents[3])
	assert.Equal(t, tails[0], fmt.Sprintf("lib/errors/combo_error_test.go:%v: code = NotFound details = {", line+4), "first line should start combined error and contain filename and line number")
	assert.Equal(t, tails[1], fmt.Sprintf("lib/errors/combo_error_test.go:%v: code = Unknown desc = Error}", line+1), "second line should contain combined external error and filename and line number")
	assert.Equal(t, tails[2], fmt.Sprintf("lib/errors/combo_error_test.go:%v: code = NotFound desc = Err1}", line+2), "third line should contain second combined error and filename and line number")
	assert.Equal(t, tails[3], fmt.Sprintf("lib/errors/combo_error_test.go:%v: code = NotFound desc = Err2}}}", line+3), "fourth line should start combined error and contain filename and line number")
	assert.Contains(t, err.NilOrError().Error(), "Error", "error should contain the error message in error chain")
	assert.Contains(t, err.NilOrError().Error(), "Err1", "error should contain the error message in combo error")
	assert.Contains(t, err.NilOrError().Error(), "Err2", "error should contain the error message in combo error")
}

// Detail(err2) will print out
// ```
// {$GOPATH/src/improbable.io/lib/errors/combo_error_test.go:124: code = Internal desc = Err3}
// {$GOPATH/src/improbable.io/lib/errors/combo_error.go:123: code = NotFound details = {
//   {$GOPATH/src/improbable.io/lib/errors/combo_error_test.go:121: code = Unknown desc = Err1}
//   {$GOPATH/src/improbable.io/lib/errors/combo_error_test.go:122: code = NotFound desc = Err2}}}
// ```
func TestComboErrWithErrChain(t *testing.T) {
	sharedflags.Set.Set("errors_include_debug_info", "true")
	e := NewCombo()
	_, _, line, _ := runtime.Caller(0) // line+0
	e.Add(fmt.Errorf("Err1"))          // line+1
	e.Add(New(nil, NotFound, "Err2"))  // line+2
	err := e.NilOrError()              // line+3
	err2 := New(err, Internal, "Err3") // line+4
	indents, tails := splitDetailsForTest(t, err2)
	assert.Equal(t, "", indents[0], "first line should be unindented")
	assert.Equal(t, "", indents[1], "second line should be unindented")
	assert.Equal(t, "\t", indents[2], "first line inside combo error should be indented once")
	assert.Equal(t, "\t", indents[3], "second line inside combo error should be indented once")
	assert.Equal(t, tails[0], fmt.Sprintf("lib/errors/combo_error_test.go:%v: code = Internal desc = Err3}", line+4), "first line should contain outer error with file and line number")
	assert.Equal(t, tails[1], fmt.Sprintf("lib/errors/combo_error_test.go:%v: code = NotFound details = {", line+3), "second line should contain outer error with file and line number")
	assert.Equal(t, tails[2], fmt.Sprintf("lib/errors/combo_error_test.go:%v: code = Unknown desc = Err1}", line+1), "third line should contain outer error with file and line number")
	assert.Equal(t, tails[3], fmt.Sprintf("lib/errors/combo_error_test.go:%v: code = NotFound desc = Err2}}}", line+2), "fourth line should contain outer error with file and line number")
	assert.NotContains(t, err2.Error(), "Err1", "error should contain the error message in combo error")
	assert.NotContains(t, err2.Error(), "Err2", "error should contain the error message in combo error")
	assert.Contains(t, err2.Error(), "Err3", "error should contain the error message in combo error")
}

// Detail(err2) will print out
// ```
// {$GOPATH/src/improbable.io/lib/errors/combo_error_test.go:157: code = Internal desc = Err3}
// {$GOPATH/src/improbable.io/lib/errors/combo_error_test.go:156: code = NotFound details = {
//   {$GOPATH/src/improbable.io/lib/errors/combo_error_test.go:154: code = Unknown desc = Err2}
//   {$GOPATH/src/improbable.io/lib/errors/combo_error_test.go:155: code = NotFound details = {
//     {$GOPATH/src/improbable.io/lib/errors/combo_error_test.go:153: code = NotFound desc = Err1}}}}}
// ```
func TestComboErrInComboErr(t *testing.T) {
	sharedflags.Set.Set("errors_include_debug_info", "true")
	c := NewCombo()
	c2 := NewCombo()

	_, _, line, _ := runtime.Caller(0) // line+0
	c.Add(New(nil, NotFound, "Err1"))  // line+1
	c2.Add(fmt.Errorf("Err2"))         // line+2
	c2.Add(c.NilOrError())             // line+3
	err := c2.NilOrError()             // line+4
	err2 := New(err, Internal, "Err3") // line+5

	indents, tails := splitDetailsForTest(t, err2)
	assert.Equal(t, "", indents[0], "first line should be unindented")
	assert.Equal(t, "", indents[1], "second line should be unindented")
	assert.Equal(t, "\t", indents[2], "first line inside combo error should be indented once")
	assert.Equal(t, "\t", indents[3], "second line inside combo error should be indented once")
	assert.Equal(t, "\t\t", indents[4], "line inside combo error inside combo error should be indented once")
	assert.Equal(t, tails[0], fmt.Sprintf("lib/errors/combo_error_test.go:%v: code = Internal desc = Err3}", line+5), "first line should contain outer error with file and line number")
	assert.Equal(t, tails[1], fmt.Sprintf("lib/errors/combo_error_test.go:%v: code = NotFound details = {", line+4), "second line should contain outer error with file and line number")
	assert.Equal(t, tails[2], fmt.Sprintf("lib/errors/combo_error_test.go:%v: code = Unknown desc = Err2}", line+2), "third line should contain outer error with file and line number")
	assert.Equal(t, tails[3], fmt.Sprintf("lib/errors/combo_error_test.go:%v: code = NotFound details = {", line+3), "fourth line should contain outer error with file and line number")
	assert.Equal(t, tails[4], fmt.Sprintf("lib/errors/combo_error_test.go:%v: code = NotFound desc = Err1}}}}}", line+1), "fifth line should contain outer error with file and line number")
	assert.NotContains(t, err2.Error(), "Err1", "error should contain the error message in combo error")
	assert.NotContains(t, err2.Error(), "Err2", "error should contain the error message in combo error")
	assert.Contains(t, err2.Error(), "Err3", "error should contain the error message in combo error")
}
