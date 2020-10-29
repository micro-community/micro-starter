#enviroments
NAME=auth
IMAGE_NAME=micro-community/$(NAME)
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_TAG=$(shell git describe --abbrev=0 --tags --always --match "v*")
GIT_IMPORT=github.com/micro-community/micro-starter/config
CGO_ENABLED=0
BUILD_DATE=$(shell date +%s)
LDFLAGS=-X $(GIT_IMPORT).BuildDate=$(BUILD_DATE) -X $(GIT_IMPORT).GitCommit=$(GIT_COMMIT) -X $(GIT_IMPORT).GitTag=$(GIT_TAG)
IMAGE_TAG=$(GIT_TAG)-$(GIT_COMMIT)

#vars
GOPATH:=$(shell go env GOPATH)
GOROOT:=$(shell go env GOROOT)
empty :=
space := $(empty) $(empty)

#go imports
VALIDATE_IMPORT := Mvalidate/validate.proto=github.com/envoyproxy/protoc-gen-validate/validate
GO_IMPORT_SPACES := ${VALIDATE_IMPORT},\
	Mgoogle/protobuf/any.proto=github.com/golang/protobuf/ptypes/any,\
	Mgoogle/protobuf/duration.proto=github.com/golang/protobuf/ptypes/duration,\
	Mgoogle/protobuf/struct.proto=github.com/golang/protobuf/ptypes/struct,\
	Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,\
	Mgoogle/protobuf/wrappers.proto=github.com/golang/protobuf/ptypes/wrappers,\
	Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor


GO_IMPORT:=$(subst $(space),,$(GO_IMPORT_SPACES))

all: build

vendor:
	go mod vendor

build:
	go build -a -installsuffix cgo -ldflags "-s -w ${LDFLAGS}" -o $(NAME)

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	# docker build . -t auth:latest
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .
	docker tag $(IMAGE_NAME):$(IMAGE_TAG) $(IMAGE_NAME):latest
	docker push $(IMAGE_NAME):$(IMAGE_TAG)
	docker push $(IMAGE_NAME):latest

# make protos

.PHONY: proto
proto:
	protoc \
	-I "${GOROOT}/include" \
	-I "protos/${p}"  \
	-I "${GOPATH}/src" \
	--go_out="${GO_IMPORT}:protos/${p}"  \
	--micro_out="${GO_IMPORT}:protos/${p}"   \
	--validate_out="lang=go:protos/${p}"   \
  "${p}.proto"


.PHONY: message
message:
	protoc  \
	-I protos/message \
	--go_out="${GO_IMPORT}:protos/message"  \
	--micro_out="${GO_IMPORT}:protos/message"   \
	--validate_out="lang=go:protos/message"   \
	message.proto
