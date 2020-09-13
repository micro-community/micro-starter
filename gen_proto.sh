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


# github.com/golang/protobuf/ptypes/any
# github.com/golang/protobuf/descriptor
#或者
#github.com/gogo/protobuf/types

MOD=Mgoogle/protobuf/any.proto=github.com/golang/protobuf/ptypes/any

#Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor
protoc -Iprotos --go_out=${MOD},plugins=grpc:protos updatecolumn.proto

# 增加valid 校验
# protoc-go-inject-tag -input=protos/updatecolumn.pb.go


# 去掉omitempty
# ls *.pb.go | xargs -n1 -IX bash -c "sed -e 's/,omitempty//' X > X.tmp && mv X{.tmp,}"

# 增加 bson:"-"
# ls *.pb.go | xargs -n1 -IX bash -c "sed -e 's/json\:\"-\"/json\:\"-\" bson:\"-\"/' X > X.tmp && mv X{.tmp,}"
# ls *.pb.go | xargs -n1 -IX bash -c "sed -e 's/json\:\"id\"/json\:\"id\" bson:\"_id\"/' X > X.tmp && mv X{.tmp,}"
# ls *.pb.go | xargs -n1 -IX bash -c "sed -e 's/json\:\"children\"/json\:\"children\" bson:\"-\"/' X > X.tmp && mv X{.tmp,}"

echo "Complete"

read -n1 -p "Press any key to continue..."