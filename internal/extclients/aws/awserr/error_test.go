// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package awserr

import (
	"testing"

	"github.com/max1314/imp-billing-datalake/lib/errors"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRemapErr_NilErr(t *testing.T) {
	remappedErr := RemapError(nil)
	require.NoError(t, remappedErr)
}

func TestRemapErr_NoAWSErr(t *testing.T) {
	err := errors.New(nil, errors.Aborted, "aborted")

	remappedErr := RemapError(err)
	require.Error(t, remappedErr)
	assert.Exactly(t, remappedErr, err)
}
func TestRemapErr_AWSInvalidParameterErr(t *testing.T) {
	err := awserr.New(string(InvalidParameterErrCode), "invalid parameter", nil)

	remappedErr := RemapError(err)
	require.Error(t, remappedErr)
	assert.Equal(t, errors.InvalidArgument, errors.Code(remappedErr))
}
