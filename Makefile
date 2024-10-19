LOCAL_BIN:=$(CURDIR)/bin
API_PROTO:=chat_api_v1
API:=chat_api

<<<<<<< HEAD
BUILD_DIR:=./build

=======
>>>>>>> 2cd7fc411283ce115961b3f8b7bd47e5b62ed01c
lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
<<<<<<< HEAD
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0
=======
>>>>>>> 2cd7fc411283ce115961b3f8b7bd47e5b62ed01c

get-deps:
	GOBIN=$(LOCAL_BIN) go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc


generate:
	make generate-chat-api

generate-chat-api:
	mkdir -p pkg/$(API)
	protoc --proto_path api/$(API_PROTO) \
	--go_out=pkg/$(API) --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/$(API) --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/$(API_PROTO)/*

db-up:
	docker compose -f ${BUILD_DIR}/docker-compose.yml up --build -d pg migrator

db-down:
	docker compose -f ${BUILD_DIR}/docker-compose.yml down pg migrator
