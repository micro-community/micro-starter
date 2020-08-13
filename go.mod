module auth-demo

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/coreos/etcd v3.3.22+incompatible // indirect
	github.com/go-log/log v0.2.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/gorm v1.9.15
	github.com/lucas-clemente/quic-go v0.17.3 // indirect
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
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/crypto v0.0.0-20200728195943-123391ffb6de // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	golang.org/x/sys v0.0.0-20200812155832-6a926be9bd1d // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/tools v0.0.0-20200812231640-9176cd30088c // indirect
	google.golang.org/genproto v0.0.0-20200812160120-2e714abc8b50 // indirect
	google.golang.org/grpc v1.31.0 // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/redis.v5 v5.2.9
	honnef.co/go/tools v0.0.1-2020.1.5 // indirect
)
