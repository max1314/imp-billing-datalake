// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package errors

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaybePanicRecover(t *testing.T) {
	var err error
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer MaybePanicRecover(&err)
		panic("panic test")
	}()
	wg.Wait()
	require.Error(t, err)
	require.Contains(t, err.Error(), "panic recover\n panic test\n stack trace")
}

func TestMaybePanicRecoverCh(t *testing.T) {
	errCh := make(chan error)
	go func() {
		defer MaybePanicRecoverCh(errCh)
		panic("panic test")
	}()
	err := <-errCh
	require.Error(t, err)
	require.Contains(t, err.Error(), "panic recover\n panic test\n stack trace")
}
