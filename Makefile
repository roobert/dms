all:
	@go build -o dms client.go  db.go  error.go  handler_prometheus.go  init.go  main.go  post_data_prometheus.go  routes.go  tables.go
run:
	@bash -c 'go run client.go  db.go  error.go  handler_prometheus.go  init.go  main.go  post_data_prometheus.go  routes.go  tables.go |& ~/go/bin/pp'
test:
	@go test -v
clean:
	@rm -v dms
