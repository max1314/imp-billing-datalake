// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package buffer

import (
	"io"
)

func NewClosingBuffer(onClose func([]byte) error) io.WriteCloser {
	return &closingBuffer{buf: []byte{}, onClose: onClose}
}

type closingBuffer struct {
	buf     []byte
	onClose func([]byte) error
}

func (cb *closingBuffer) Write(b []byte) (n int, err error) {
	cb.buf = append(cb.buf, b...)
	return len(b), nil
}

func (cb *closingBuffer) Close() error {
	return cb.onClose(cb.buf)
}
