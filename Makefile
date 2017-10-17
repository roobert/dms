SERVER_FILES := $(filter-out server/*_test.go, $(wildcard server/*.go))

all: build-server

clean:
	@rm -v dms

test:
	@go test ./...

build-server:
	@go build -o dms ${SERVER_FILES}

build-client:
	@go build -o dms-handler

run-server:
	@go run ${SERVER_FILES}
