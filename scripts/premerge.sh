#!/bin/bash

# Copyright (c) Improbable Worlds Ltd, All Rights Reserved

# Get build + test dependencies. -d also doesn't bother with installing the
# packages, it just downloads them
#circle local presubmit -m check
go fmt

go vet

export S3_FORCE_PATH_STYLE=true
go get -t -d
go test