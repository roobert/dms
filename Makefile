all: build-server build-client

SERVER_FILES := $(filter-out server/*_test.go, $(wildcard server/*.go))


build-server:
	@go build -o dms ${SERVER_FILES}

test:
	@go test ./...

build-client:
	@go build -o dms-handler ls handler/*.go -I handler/*_test.go
run-server:
	@go run server/*.go -I server/*_test.go
clean:
	@rm -v dms
