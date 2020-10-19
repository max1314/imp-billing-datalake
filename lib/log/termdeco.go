// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package log

import (
	"io"
	"io/ioutil"

	"github.com/tatsushid/termdeco"
)

type termdecoWriter struct {
	w io.Writer
}

func (w termdecoWriter) Write(p []byte) (n int, err error) {
	n, err = termdeco.Fprintf(w.w, "%s", p)
	if err != nil {
		return n, err
	}
	// termdeco may write more than len(p) bytes, causing an error in io.Copy. On success, return
	// the original length to avoid this
	return len(p), nil
}

func (w termdecoWriter) ReadFrom(r io.Reader) (n int64, err error) {
	b, err := ioutil.ReadAll(r)
	n = int64(len(b))
	if err != nil {
		return n, err
	}
	i, err := termdeco.Fprint(w.w, b)
	if err != nil {
		return int64(i), err
	}
	// termdeco may write more than len(p) bytes, causing an error in io.Copy. On success, return
	// the original length to avoid this
	return n, nil
}

var _ io.ReaderFrom = termdecoWriter{}
var _ io.Writer = termdecoWriter{}
