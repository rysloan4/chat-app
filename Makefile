run-chat:
	go run main.go hub.go conn.go

fmt:
	gofmt -l -s -w .

lint:
	golint ./...
