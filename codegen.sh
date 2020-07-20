#!/bin/bash
set -e

if [ -f ./build ]; then
	find ./build -name gen-* -delete
fi

if [ -f ./.env ]; then
  source .env
fi;

_PWD=$PWD

function yellow {
	echo -e "\033[33m$@\033[39m"
}
function green {
	echo -e "\033[32m$@\033[39m"
}

function gofmt {
	yellow "> go fmt ./..."
	go fmt ./...
	green "OK"
}

function provision {
	yellow "> provision files"
	for FOLDER in system compose messaging; do
   	$GOPATH/bin/statik -p $FOLDER -m -Z -f -src="./provision/$FOLDER/src" -dest "./provision"
	done
	green "OK"
}


function events {
  if [ ! -f "build/event-gen" ]; then
		CGO_ENABLED=0 go build -o ./build/event-gen ./codegen/v2/events
	fi

	for SERVICE in system compose messaging; do
	  yellow "> event files for ${SERVICE}"
	  ./build/event-gen --service ${SERVICE}
	done
	green "OK"
}


function specs {
	yellow "> specs"
	if [ ! -f "build/gen-spec" ]; then
		CGO_ENABLED=0 go build -o ./build/gen-spec codegen/v2/spec.go
	fi
	_PWD=$PWD
	SPECS=$(find $PWD -name 'spec.json' | xargs -n1 dirname)
	for SPEC in $SPECS; do
		yellow "> spec $SPEC"
		cd $SPEC && rm -rf spec && ../../build/gen-spec && cd $_PWD
		green "OK"
	done

	for SPEC in $SPECS; do
		SRC=$(basename $SPEC)
		if [ -d "codegen/$SRC" ]; then
			yellow "> README $SRC"
			codegen/codegen.php $SRC
			rsync -a codegen/common/ $SRC/
			green "OK"
		fi
	done
}


function proto {
	yellow "> proto"

	# Where should we look for the files
	PROTOBUF_PATH="codegen/corteza-protobuf"
	CORTEZA_PROTOBUF_PATH=${CORTEZA_PROTOBUF_PATH:-"${PROTOBUF_PATH}"}

  # Download protobufs to the primary location
  BRANCH=${BRANCH:-"develop"}
  ZIP="${BRANCH}.zip"
  URL=${URL:-"https://github.com/cortezaproject/corteza-protobuf/archive/${ZIP}"}
  rm -rf "${PROTOBUF_PATH}"
  curl -s --location "${URL}" > "codegen/${ZIP}"
  unzip -qq -o -d "codegen/" "codegen/${ZIP}"
  mv -f "codegen/corteza-protobuf-${BRANCH}" "${PROTOBUF_PATH}"

  DIR=./pkg/corredor
  mkdir -p ${DIR}
	yellow "  ${CORTEZA_PROTOBUF_PATH} >> ${DIR}"
	PATH=$PATH:$GOPATH/bin protoc \
		--proto_path ${CORTEZA_PROTOBUF_PATH} \
		--go_out="plugins=grpc:./${DIR}" \
		service-corredor.proto

	yellow "  ${CORTEZA_PROTOBUF_PATH} >> system/proto"
	PATH=$PATH:$GOPATH/bin protoc \
		--proto_path ${CORTEZA_PROTOBUF_PATH}/system \
		--go_out=plugins=grpc:system/proto \
		user.proto role.proto
  green "OK"
}

case ${1:-"all"} in
  provision)
    provision
    ;;
  specs)
    specs
    ;;
  proto)
    proto
    ;;
  events)
    events
    ;;
  all)
    types
    database
    provision
    specs
    proto
    events
esac

# Always finish with fmt
gofmt
