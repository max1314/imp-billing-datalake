module imp-billing-datalake

replace improbable.io/lib/db/blob => ./src/improbable.io/lib/db/blob

replace improbable.io/lib/db/blob/gcs => ./src/improbable.io/lib/db/blob/gcs

replace improbable.io/lib/db/blob/s3 => ./src/improbable.io/lib/db/blob/s3

replace improbable.io/lib/db/blob/util => ./src/improbable.io/lib/db/blob/util

replace improbable.io/lib/db/hash => ./src/improbable.io/lib/db/blob/hash

replace improbable.io/lib/errors => ./src/improbable.io/lib/errors

replace improbable.io/lib/sharedflags => ./src/improbable.io/lib/sharedflags

replace improbable.io/lib/extclients/aws/awserr => ./src/improbable.io/lib/extclients/aws/awserr

replace improbable.io/lib/testing/flagtest => ./src/improbable.io/lib/testing/flagtest

go 1.15

require (
	github.com/google/go-cmp v0.5.2 // indirect
	github.com/mwitkow/go-flagz v0.0.0-20170404101818-12def25b6a92 // indirect
	github.com/prometheus/client_golang v1.8.0 // indirect
	github.com/spf13/cobra v1.1.0 // indirect
	golang.org/x/net v0.0.0-20200904194848-62affa334b73 // indirect
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208 // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20201015140912-32ed001d685c // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	improbable.io/lib/db/blob v0.0.0-00010101000000-000000000000
	improbable.io/lib/errors v0.0.0-00010101000000-000000000000 // indirect
	improbable.io/lib/sharedflags v0.0.0-00010101000000-000000000000 // indirect
	improbable.io/lib/testing/flagtest v0.0.0-00010101000000-000000000000 // indirect
)
