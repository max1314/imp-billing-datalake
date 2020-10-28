// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package buffer

import (
	"fmt"
	"io"

	"github.com/improbable/imp-billing-datalake/errors"
)

const MaxBufferSize = 500 * 1024 * 1024 // 500M bytes

func NewClosingBuffer(onClose func([]byte) error) io.WriteCloser {
	return &closingBuffer{buf: []byte{}, onClose: onClose}
}

type closingBuffer struct {
	buf     []byte
	onClose func([]byte) error
}

func (cb *closingBuffer) Write(b []byte) (n int, err error) {
	size := len(cb.buf) + len(b)
	if size > MaxBufferSize {
		return 0, errors.New(nil, fmt.Sprintf("buffer size over max limitation: %v MB", MaxBufferSize/2014/1024))
	}
	cb.buf = append(cb.buf, b...)
	return size, nil
}

func (cb *closingBuffer) Close() error {
	return cb.onClose(cb.buf)
}
