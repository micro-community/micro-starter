module github.com/crazybber/user

go 1.15

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/go-log/log v0.2.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/gorm v1.9.16
	github.com/lucas-clemente/quic-go v0.18.0 // indirect
	github.com/marten-seemann/qtls v0.10.0 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v3 v3.0.0-alpha.0.20200731140532-31ed4aa0e8d3
	github.com/micro/go-plugins v1.5.1
	github.com/micro/micro/v3 v3.0.0-alpha
	github.com/miekg/dns v1.1.31 // indirect
	github.com/nats-io/jwt v1.0.1 // indirect
	github.com/nats-io/nats.go v1.10.0 // indirect
	github.com/nats-io/nkeys v0.2.0 // indirect
	github.com/opentracing/opentracing-go v1.1.0
	github.com/prometheus/common v0.10.0
	github.com/spf13/cobra v1.0.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/sys v0.0.0-20200831180312-196b9ba8737a // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/tools v0.0.0-20200903005429-2364a5e8fdcf // indirect
	google.golang.org/genproto v0.0.0-20200903010400-9bfcb5116336 // indirect
	google.golang.org/grpc v1.31.0 // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/redis.v5 v5.2.9
	honnef.co/go/tools v0.0.1-2020.1.5 // indirect
)
