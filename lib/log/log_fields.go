// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

// Extract logrus.Fields from proto messages and other suitably tagged objects.

package log

import (
	"crypto/sha256"
	"encoding/hex"

	"imp-billing-datalake/lib/errors"
)

type TaggedObject interface {
	// LogFields generates fields that should be attached to any log message produced while processing the object.
	LogFields() map[string]string
}

// LogFields converts the fields from a TaggedObject to logrus.Fields.
func LogFields(object TaggedObject) Fields {
	fields := make(Fields)
	for k, v := range object.LogFields() {
		fields[k] = v
	}
	return fields
}

// ExtractLogFields extracts logrus.Fields from a type-erased TaggedObject.
func ExtractLogFields(object interface{}) Fields {
	tagged, ok := object.(TaggedObject)
	if !ok {
		return make(Fields)
	}
	return LogFields(tagged)
}

// MakeSecret one-way hashes values (unless they are unset).
// This is useful since we want to not store personally-identifiable information where unnecessary
// but we _do_ want to still be able to correlate on unique values.
func MakeSecret(input string, unsetValue string) (string, error) {
	if input == unsetValue {
		return unsetValue, nil
	}
	hash := sha256.New()
	_, err := hash.Write([]byte(input))
	if err != nil {
		return "", errors.Wrapf(err, "failed to sha256 hash")
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
