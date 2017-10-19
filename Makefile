LISTENER_FILES := $(filter-out cmd/listener/*_test.go, $(wildcard cmd/listener/*.go))
HANDLER_FILES := $(filter-out cmd/handler/*_test.go, $(wildcard cmd/handler/*.go))

all: dms-listener dms-handler

clean:
	@rm -v dms-handler
	@rm -v dms-listener

test:
	@go test ./...

dms-listener:
	@go build -o dms-listener ${LISTENER_FILES}

dms-handler:
	@go build -o dms-handler ${HANDLER_FILES}

run-listener:
	@go run ${LISTENER_FILES}

run-handler:
	@go run ${HANDLER_FILES}
