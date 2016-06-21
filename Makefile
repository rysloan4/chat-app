run-chat:
	go build main.go hub.go conn.go
	./main

fmt:
	gofmt -l -s -w .

lint:
	golint ./...

build:
    go build main.go hub.go conn.go