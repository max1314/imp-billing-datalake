// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

// package debug serves as an easy way to include flags that will make debugging servers locally more pleasant.
// This can also be helpful when writing tests so you know where errors are originating from.
//
// Usage
//
// Add the following line to any file that will be included in the build
//	import _ "improbable.io/lib/errors/include_debug"
package debug

import "improbable.io/lib/sharedflags"

func init() {
	sharedflags.Set.Set("errors_include_debug_info", "true")
	sharedflags.Set.Set("log_level", "debug")
}
