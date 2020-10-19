// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubTaggedStruct struct {
	fields map[string]string
}

func (s *stubTaggedStruct) LogFields() map[string]string {
	return s.fields
}

func TestExtractFieldsFromTaggedStruct(t *testing.T) {
	require.Implements(t, (*TaggedObject)(nil), &stubTaggedStruct{})
	fields := ExtractLogFields(&stubTaggedStruct{fields: map[string]string{
		"field": "a",
	}})
	assert.Equal(t, Fields{"field": "a"}, fields)
}

func TestMakeSecret_WhenASecretIsUnset_ShouldNotHash(t *testing.T) {
	unset := "unset"
	result, err := MakeSecret(unset, unset)
	assert.NoError(t, err)
	assert.Equal(t, unset, result, "No point hashing it if it's unset, just say it's unset.")
}

func TestMakeSecret_WhenASecretIsSet_ShouldHash(t *testing.T) {
	unset := "[unset]"
	result, err := MakeSecret("foobar", unset)
	assert.NoError(t, err)
	assert.NotEqual(t, unset, result, "Should definitely not be the unset string")
	assert.NotEqual(t, "foobar", result, "Should definitely not come through unobscured")
	assert.Equal(t, "c3ab8ff13720e8ad9047dd39466b3c8974e592c2fa383d4a3960714caef0c4f2", result, "Should be the hex-encoding of the hashed value (filesystem and URL safe)")
}
