// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package blob

import (
	"io"
	"io/ioutil"
	"mime"
	"path/filepath"

	"imp-billing-datalake/lib/errors"
)

// ReadAll will read the Contents of the FakeBlob at the supplied path in a given
// bucket or return an error if the FakeBlob does not exist or could not be read.
func ReadAll(b BucketReader, p string) ([]byte, error) {
	r, err := b.NewReader(p, nil)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	res, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.Newf(err, errors.Internal, "error reading FakeBlob: %s. %v", p, err)
	}
	return res, nil
}

func NewReader(b BucketReader, p string) (io.ReadCloser, error) {
	return b.NewReader(p, nil)
}

func GuessContentType(path string) string {
	ret := mime.TypeByExtension(filepath.Ext(path))
	if ret == "" {
		ret = "application/octet-stream"
	}
	return ret
}
