add-submodule:
	git submodule add git@github.com:mf-sakura/bh_proto.git proto

proto-gen-go:
	protoc --proto_path=proto/user/v1/ --go_out=plugins=grpc:app/proto proto/user/v1/user.proto

update-submodule:
	git submodule foreach 'git fetch;git checkout master; git pull'

db-test:
	go test github.com/mf-sakura/bh_user/app/db

build:
	go build -o bh_user github.com/mf-sakura/bh_user/app/cmd