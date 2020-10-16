// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

// utilities for converting artifact md5 hash to different form of encodings

package hash

import (
	"encoding/base64"
	"encoding/hex"

	"improbable.io/lib/errors"
)

// Md5 is a type that stores the byte representation of a md5 checksum.
type Md5 []byte

// NewMd5FromHexString converts the hex string version of a checksum to its byte Md5 representation.
func NewMd5FromHexString(hexMd5 string) (Md5, error) {
	md5, err := hex.DecodeString(hexMd5)
	if err != nil {
		return nil, errors.Newf(err, errors.InvalidArgument, "invalid md5: %v", err)
	}
	return md5, nil
}

// NewMd5FromBase64String converts the base64 string version of a checksum to its byte Md5 representation.
func NewMd5FromBase64String(base64Md5 string) (Md5, error) {
	md5, err := base64.StdEncoding.DecodeString(base64Md5)
	if err != nil {
		return nil, errors.Newf(err, errors.InvalidArgument, "invalid md5: %v", err)
	}
	return md5, nil
}

// ToString converts the byte Md5 representation to its hex representation.
func (md5 Md5) ToString() string {
	return hex.EncodeToString(md5)
}

// ToBase64String converts the byte Md5 representation to its base64 string representation.
func (md5 Md5) ToBase64String() string {
	return base64.StdEncoding.EncodeToString(md5)
}
