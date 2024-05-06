.PHONY: build_proto_go images

all: build_proto_go images

build_proto_go:
	 protoc --go_out=protobufs --go_opt=paths=source_relative --go-grpc_out=protobufs --go-grpc_opt=paths=source_relative protobufs/apiserver.proto

images:
	cd images && ./build-images.sh