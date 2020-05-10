
GOPATH:=$(shell go env GOPATH)
MODIFY=Mgithub.com/micro/go-micro/api/proto/api.proto=github.com/micro/go-micro/v2/api/proto

.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/user/user.proto
	go-bindata -pkg migrations -ignore bindata -prefix ./datastore/migrations/ -o ./datastore/migrations/bindata.go ./datastore/migrations
    

.PHONY: build
build: proto
	go-bindata -pkg migrations -ignore bindata -prefix ./datastore/migrations/ -o ./datastore/migrations/bindata.go ./datastore/migrations
	go build -o user-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t user-service:latest

migrate:
	go-bindata -pkg migrations -ignore bindata -prefix ./datastore/migrations/ -o ./datastore/migrations/bindata.go ./datastore/migrations

table:
	migrate create -ext sql -dir ./datastore/migrations -seq create_users

install:
	go get \
		github.com/golang/protobuf/protoc-gen-go \
		github.com/tmthrgd/go-bindata/... \
		github.com/golang/mock/mockgen
