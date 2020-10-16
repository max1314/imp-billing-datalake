// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package errors_test

import (
	"fmt"
	"testing"

	"improbable.io/lib/errors"

	"github.com/stretchr/testify/assert"
)

func TestCauseMapping(t *testing.T) {

	for _, spec := range []struct {
		candidate error
		expected  errors.Category
	}{
		{
			candidate: errors.New(nil, errors.NotFound, "foo bar"),
			expected:  errors.User,
		},
		{
			candidate: errors.New(nil, errors.InvalidArgument, "something invalid"),
			expected:  errors.User,
		},
		{
			candidate: errors.New(nil, errors.Unknown, "something unknown"),
			expected:  errors.System,
		},
		{
			candidate: errors.New(nil, errors.Unavailable, "unavailable"),
			expected:  errors.Transient,
		},
		{
			candidate: fmt.Errorf("undefined error is Unknown by default"),
			expected:  errors.System,
		},
	} {
		assert.Equal(t, spec.expected, errors.Categorise(spec.candidate))
	}
}
