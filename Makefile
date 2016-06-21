run-chat:
	make build && ./main

fmt:
	gofmt -l -s -w .

lint:
	golint ./...

build:
	go build main.go hub.go conn.go
