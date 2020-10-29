module github.com/micro-community/micro-starter

go 1.15

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.29.0

require (
	github.com/dgraph-io/dgo/v200 v200.0.0-20201023081658-a9ad93fe6ebd
	github.com/envoyproxy/protoc-gen-validate v0.4.1
	github.com/go-redis/redis/v8 v8.3.2
	github.com/golang/protobuf v1.4.3
	github.com/gomodule/redigo/redis v0.0.0-20200429221454-e14091dffc1b
	github.com/hashicorp/go-version v1.2.1
	github.com/lib/pq v1.8.0
	github.com/micro/micro/v3 v3.0.0-beta.7
	github.com/olivere/elastic/v7 v7.0.21
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.7.0
	github.com/urfave/cli/v2 v2.2.0
	go.mongodb.org/mongo-driver v1.4.2
	go.uber.org/dig v1.10.0
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9
	google.golang.org/grpc v1.33.1
	google.golang.org/protobuf v1.25.0
	gopkg.in/sohlich/elogrus.v7 v7.0.0
	gorm.io/driver/mysql v1.0.3
	gorm.io/driver/sqlite v1.1.3
	gorm.io/gorm v1.20.5
)
