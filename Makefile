all:
	@go build -o dms bitmap.go  client.go  db.go  error.go  handler_prometheus.go  init.go  main.go  routes.go
run:
	@bash -c 'go run bitmap.go  client.go  db.go  error.go  handler_prometheus.go  init.go  main.go  routes.go |& ~/go/bin/pp'
test:
	@go test -v
clean:
	@rm -v dms
