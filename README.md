# micro starter

micro v3 starter framework

micro(v3) 单(微)服务能力集框架,有兴趣的朋友，热烈欢迎 PR.

## install

```bash
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
go get github.com/micro/micro/v3/cmd/protoc-gen-micro
```

```bash
protoc -I.  --go_out=protos --micro_out=protos protos/user.proto
protoc -Iprotos  --go_out=protos --micro_out=protos role.proto
```

### 异步事件

```bash
protoc -Iprotos  --go_out=protos/message --micro_out=protos/message message/message.proto
```

### the following for graph design for dgraph

```bash
protoc -Iprotos/rbac  -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate  --go_out=protos/rbac --micro_out=protos/rbac --validate_out="lang=go:protos/rbac" rbac.proto
```

## Coding Style

- 尽量给全部的代码加上注释，关键函数和类一定要加上注释
- 代码的文件夹结构尽量展平，不要多层嵌套
- 文件夹名称全部小写，不使用驼峰
- 尽量避免长文件和文件夹名称
- 文件名有多个单词时，使用下划线连接
- 单词拼写使用 cspell 检查
- 面向 TDD 开发，关键的 UT 要写

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service

```
./auth
```

Build a docker image

```
make docker
```
