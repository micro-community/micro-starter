# 说明

```proto
import "google/protobuf/descriptor.proto";
import "google/protobuf/duration.proto";
import "google/protobuf/timestamp.proto";
```

```shell
protoc -I protos/rbac  -I $env:GOPATH/src  --go_out=protos/rbac --micro_out=protos/rbac --validate_out="lang=go:protos/rbac" rbac.proto
```


```shell
protoc -I protos/rbac  -I ${go env GOROOT} -I ${go env GOPATH}/src  --go_out=protos/rbac --micro_out=protos/rbac --validate_out="lang=go:protos/rbac" rbac.proto
```
