# micro starter

micro v3 starter framework

micro(v3) 微服务能力集框架,有兴趣的朋友，热烈欢迎 PR.

## install

```bash
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
go get github.com/micro/micro/v3/cmd/protoc-gen-micro
```

### generate protos

- `make user`
- `make role`
- `make message` #async message

### following for graph design for dgraph

`make rbac`

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
