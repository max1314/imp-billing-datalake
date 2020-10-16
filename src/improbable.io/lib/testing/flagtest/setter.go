// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

/*

This package helps avoid tests making flag changes and forgetting to reset them.

If you use this package, first consider changing your test and/or APIs to not require flag changes
in the first place.

Example:

	func TestSomething(t *testing.T) {
		defer flagtest.Bool(fSomeFlag, true)()
	}

*/

package flagtest

import (
	"time"
)

// Bool lets a bool flag be set for the duration of a test.
// defer the returned callback to ensure the value is reset.
// e.g.: `defer flagtest.Bool(fMyFlag, true)()`
func Bool(flag *bool, value bool) func() {
	original := *flag
	*flag = value
	return func() { *flag = original }
}

// Duration lets a duration flag be set for the duration of a test.
// defer the returned callback to ensure the value is reset.
// e.g.: `defer flagtest.Duration(fMyFlag, time.Minute)()`
func Duration(flag *time.Duration, value time.Duration) func() {
	original := *flag
	*flag = value
	return func() { *flag = original }
}

// Int lets an int flag be set for the duration of a test.
// defer the returned callback to ensure the value is reset.
// e.g.: `defer flagtest.Int(fMyFlag, 42)()`
func Int(flag *int, value int) func() {
	original := *flag
	*flag = value
	return func() { *flag = original }
}

// String lets a string flag be set for the duration of a test.
// defer the returned callback to ensure the value is reset.
// e.g.: `defer flagtest.String(fMyFlag, "dishwasher")()`
func String(flag *string, value string) func() {
	original := *flag
	*flag = value
	return func() { *flag = original }
}

// StringSlice lets a string slice flag be set for the duration of a test.
// defer the returned callback to ensure the value is reset.
// e.g.: `defer flagtest.StringSlice(fMyFlag, []string{"yes", "no", "file_not_found"})()`
func StringSlice(flag *[]string, value []string) func() {
	original := *flag
	*flag = value
	return func() { *flag = original }
}
