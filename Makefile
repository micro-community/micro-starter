
NAME=auth
IMAGE_NAME=micro-community/$(NAME)
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_TAG=$(shell git describe --abbrev=0 --tags --always --match "v*")
#GIT_IMPORT=github.com/micro-community/auth
CGO_ENABLED=0
BUILD_DATE=$(shell date +%s)
#LDFLAGS=-X $(GIT_IMPORT).BuildDate=$(BUILD_DATE) -X $(GIT_IMPORT).GitCommit=$(GIT_COMMIT) -X $(GIT_IMPORT).GitTag=$(GIT_TAG)
IMAGE_TAG=$(GIT_TAG)-$(GIT_COMMIT)

empty :=
space := $(empty) $(empty)

GOPATH:=$(shell go env GOPATH)
GOROOT:=$(shell go env GOROOT)

#for go imports
VALIDATE_IMPORT := Mvalidate/validate.proto=github.com/envoyproxy/protoc-gen-validate/validate
GO_IMPORT_SPACES := ${VALIDATE_IMPORT},\
	Mgoogle/protobuf/any.proto=github.com/golang/protobuf/ptypes/any,\
	Mgoogle/protobuf/duration.proto=github.com/golang/protobuf/ptypes/duration,\
	Mgoogle/protobuf/struct.proto=github.com/golang/protobuf/ptypes/struct,\
	Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,\
	Mgoogle/protobuf/wrappers.proto=github.com/golang/protobuf/ptypes/wrappers,\
	Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor

GO_IMPORT:=$(subst $(space),,$(GO_IMPORT_SPACES))

.PHONY: testvar
testvar:
	echo ${GO_IMPORT}

.PHONY: rbac
rbac:
	protoc \
	-I "${GOROOT}/include" \
	-I protos/rbac  \
	-I "${GOPATH}/src" \
	--go_out="${GO_IMPORT}:protos/rbac"  \
	--micro_out="${GO_IMPORT}:protos/rbac"   \
	--validate_out="lang=go:protos/rbac"   \
  rbac.proto

.PHONY: user
user:
	protoc \
	-I "${GOROOT}/include" \
	-I protos  \
	-I "${GOPATH}/src" \
	--go_out="${GO_IMPORT}:protos"  \
	--micro_out="${GO_IMPORT}:protos"   \
	--validate_out="lang=go:protos"   \
  user.proto

.PHONY: role
role:
	protoc \
	-I "${GOROOT}/include" \
	-I protos \
	-I "${GOPATH}/src" \
	--go_out="${GO_IMPORT}:protos"  \
	--micro_out="${GO_IMPORT}:protos"   \
	--validate_out="lang=go:protos"   \
  role.proto

.PHONY: build
build:
	go build -o auth *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t auth:latest
