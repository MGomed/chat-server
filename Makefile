LOCAL_BIN:=$(CURDIR)/bin
API_PROTO:=chat_api_v1
API:=chat_api

BUILD_DIR:=./build

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

.PHONY: test
test:
	mkdir -p out/coverage
	go clean -testcache
	go test -coverprofile out/coverage/cover.out ./...
	go tool cover -html=out/coverage/cover.out -o out/coverage/coverage.html

install-deps:
	[ -f $(LOCAL_BIN)/golangci-lint ] || GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
	[ -f $(LOCAL_BIN)/protoc-gen-go ] || GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	[ -f $(LOCAL_BIN)/protoc-gen-go-grpc ] || GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	[ -f $(LOCAL_BIN)/goose ] || GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0
	[ -f $(LOCAL_BIN)/mockgen ] || GOBIN=$(LOCAL_BIN) go install github.com/golang/mock/mockgen@v1.6.0
	[ -f $(LOCAL_BIN)/protoc-gen-validate ] || GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v1.0.4

get-deps:
	GOBIN=$(LOCAL_BIN) go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-chat-api
	make generate_mocks

generate_mocks:
	./generate.sh

generate-chat-api:
	mkdir -p pkg/$(API)
	protoc --proto_path api/$(API_PROTO) --proto_path vendor.protogen \
	--go_out=pkg/$(API) --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/$(API) --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--validate_out lang=go:pkg/$(API) --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=bin/protoc-gen-validate \
	api/$(API_PROTO)/*

vendor-proto:
	@if [ ! -d vendor.protogen/validate ]; then \
		mkdir -p vendor.protogen/validate && \
		git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate && \
		mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate && \
		rm -rf vendor.protogen/protoc-gen-validate ; \
	fi

db-up:
	docker compose -f ${BUILD_DIR}/docker-compose.yml up --build -d pg migrator

db-down:
	docker compose -f ${BUILD_DIR}/docker-compose.yml down pg migrator
