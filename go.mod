module github.com/micro-community/auth

go 1.15

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/dgraph-io/dgo/v200 v200.0.0-20200916081436-9ff368ad829a
	github.com/envoyproxy/protoc-gen-validate v0.4.1
	github.com/go-redis/redis/v8 v8.3.1
	github.com/golang/protobuf v1.4.2
	github.com/gomodule/redigo/redis v0.0.0-20200429221454-e14091dffc1b
	github.com/hashicorp/go-version v1.2.1
	github.com/micro/go-micro/v3 v3.0.0-beta.3.0.20201009122815-dad05be95ee0
	github.com/micro/micro/v3 v3.0.0-beta.6
	github.com/olivere/elastic/v7 v7.0.20
	github.com/sirupsen/logrus v1.7.0
	github.com/urfave/cli/v2 v2.2.0
	go.mongodb.org/mongo-driver v1.4.2
	go.uber.org/dig v1.10.0
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20201012173705-84dcc777aaee
	golang.org/x/sync v0.0.0-20201008141435-b3e1573b7520
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/sohlich/elogrus.v7 v7.0.0
	gorm.io/driver/mysql v1.0.2
	gorm.io/driver/sqlite v1.1.3
	gorm.io/gorm v1.20.2
)
