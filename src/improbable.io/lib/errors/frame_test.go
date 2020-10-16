// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package errors

import (
	"strings"
	"testing"

	"improbable.io/lib/testing/flagtest"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/status"
)

func TestFrame_Error(t *testing.T) {
	assert.Equal(t, "code = Unavailable desc = foo", newFrame(nil, Unavailable, "foo").Error())
}

func TestFrame_FileAndLine(t *testing.T) {
	defer flagtest.Bool(fEnableFilenameAndLineNumber, true)()
	f, l := newFrame(nil, Unavailable, "foo").FileAndLine()
	assert.Equal(t, "testing.go", f[strings.LastIndex(f, "testing.go"):])
	assert.NotZero(t, l)
}

func TestFrame_Code(t *testing.T) {
	assert.Equal(t, Unavailable, newFrame(nil, Unavailable, "foo").Code())
}

func TestFrame_ToGRPC(t *testing.T) {
	assert.Equal(t, status.Errorf(Unavailable, "foo"), newFrame(nil, Unavailable, "foo").ToGRPC())
}
