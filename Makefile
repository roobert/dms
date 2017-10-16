all: test build-server

clean:
	@rm -v dms

test:
	@go test ./...

SERVER_FILES := $(filter-out server/*_test.go, $(wildcard server/*.go))

build-server:
	@go build -o dms ${SERVER_FILES}

build-client:
	@go build -o dms-handler

run-server:
	@go run ${SERVER_FILES}
