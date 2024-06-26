BINARY_NAME=main

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## dev: dev the application
.PHONY: dev
dev:
	air \
		--build.exclude_dir "" \
		--build.include_ext "go" \
		--misc.clean_on_exit "true" \
		--build.cmd "make build" --build.bin "./${BINARY_NAME}" --build.delay "100"

## build/proto: build the proto files
.PHONY: build/proto
build/proto:
	rm -f package/proto/*.go
	protoc \
		--go_out=package/proto \
		--go_opt=paths=source_relative \
		--go-grpc_out=package/proto \
		--go-grpc_opt=paths=source_relative \
		--proto_path=package/proto \
		package/proto/*.proto

## build: build the application
.PHONY: build
build:
	go build -o ${BINARY_NAME} main.go

## test: test the application
.PHONY: test
test:
	mkdir -p .coverage
	go test -v -race -cover -p 1 -coverprofile=.coverage/coverage.out ./...
	go tool cover -html=.coverage/coverage.out -o .coverage/coverage.html

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## clean: format code and tidy modfile
.PHONY: clean
clean:
	go clean
	rm -f ${BINARY_NAME}
	rm -f package/proto/*.go
