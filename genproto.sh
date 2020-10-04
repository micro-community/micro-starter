#!/bin/bash


set -eu

function trap_handler {
  MYSELF="$0"   # equals to my script name
  LASTLINE="$1" # argument 1: last line of error occurence
  LASTERR="$2"  # argument 2: error code of last command
  echo "Error: line ${LASTLINE} - exit status of last command: ${LASTERR}"
  exit $2
}
trap 'trap_handler ${LINENO} ${$?}' ERR


echo "Checking tools dependencies..."
which protoc
which protoc-gen-go
which protoc-gen-validate


GO_IMPORT_SPACES=Mgoogle/protobuf/any.proto=github.com/golang/protobuf/ptypes/any,\
	Mgoogle/protobuf/duration.proto=github.com/golang/protobuf/ptypes/duration,\
	Mgoogle/protobuf/struct.proto=github.com/golang/protobuf/ptypes/struct,\
	Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,\
	Mgoogle/protobuf/wrappers.proto=github.com/golang/protobuf/ptypes/wrappers,\
	Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor


#MOD=Mgoogle/protobuf/any.proto=github.com/golang/protobuf/ptypes/any
#Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor
protoc -Iprotos --go_out=${GO_IMPORT_SPACES},plugins=grpc:protos user.proto

protoc -I "${GOROOT}/include" -I protos/rbac 	-I "${GOPATH}/src" --go_out="${GO_IMPORT}:protos/rbac"--micro_out="${GO_IMPORT}:protos/rbac" --validate_out="lang=go:protos/rbac" rbac.proto

# 增加valid 校验
# protoc-go-inject-tag -input=protos/user.pb.go


# 去掉omitempty
# ls *.pb.go | xargs -n1 -IX bash -c "sed -e 's/,omitempty//' X > X.tmp && mv X{.tmp,}"

# 增加 bson:"-"
# ls *.pb.go | xargs -n1 -IX bash -c "sed -e 's/json\:\"-\"/json\:\"-\" bson:\"-\"/' X > X.tmp && mv X{.tmp,}"
# ls *.pb.go | xargs -n1 -IX bash -c "sed -e 's/json\:\"id\"/json\:\"id\" bson:\"_id\"/' X > X.tmp && mv X{.tmp,}"
# ls *.pb.go | xargs -n1 -IX bash -c "sed -e 's/json\:\"children\"/json\:\"children\" bson:\"-\"/' X > X.tmp && mv X{.tmp,}"

echo "Complete"

read -n1 -p "Press any key to continue..."
