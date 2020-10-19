module imp-billing-datalake

go 1.15

require (
	cloud.google.com/go/storage v1.0.0
	github.com/aws/aws-sdk-go v1.35.9
	github.com/fortytw2/leaktest v1.3.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/improbable-eng/go-httpwares v0.0.0-20200609095714-edc8019f93cc
	github.com/jpillora/backoff v1.0.0
	github.com/mattn/go-isatty v0.0.3
	github.com/mwitkow/go-flagz v0.0.0-20170404101818-12def25b6a92
	github.com/prometheus/client_golang v0.9.3
	github.com/prometheus/client_model v0.0.0-20190812154241-14fe0d1b01d4
	github.com/sirupsen/logrus v1.2.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.3.0
	github.com/tatsushid/termdeco v0.0.0-20160605062857-dabb2eee98f0
	golang.org/x/crypto v0.0.0-20190605123033-f99c8df09eb5
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	google.golang.org/api v0.13.0
	google.golang.org/grpc v1.28.1
)

replace (
	// not sure when google.golang.org/grpc will down.
	//github.com/grpc/grpc-go.git v1.28.1 => google.golang.org/grpc v1.28.1
)
