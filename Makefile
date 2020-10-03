
GOPATH:=$(shell go env GOPATH)
GOROOT:=$(shell go env GOROOT)

VALIDATE_IMPORT := Mvalidate/validate.proto=github.com/envoyproxy/protoc-gen-validate/validate
GO_IMPORT_SPACES := ${VALIDATE_IMPORT},\
	Mgoogle/protobuf/any.proto=github.com/golang/protobuf/ptypes/any,\
	Mgoogle/protobuf/duration.proto=github.com/golang/protobuf/ptypes/duration,\
	Mgoogle/protobuf/struct.proto=github.com/golang/protobuf/ptypes/struct,\
	Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,\
	Mgoogle/protobuf/wrappers.proto=github.com/golang/protobuf/ptypes/wrappers,\
	Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor


.PHONY: proto
proto:
	protoc \
	-I "${GOROOT}/include" \
	-I protos/rbac  \
	-I "${GOPATH}/src" \
	--go_out=protos/rbac  \
	--micro_out=protos/rbac   \
	--validate_out="lang=go:protos/rbac"   \
  rbac.proto

.PHONY: build
build:
	go build -o auth *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t auth:latest
