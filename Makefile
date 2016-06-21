GO_BUILD_ENV := GOOS=linux GOARCH=amd64
DOCKER_BUILD=$(shell pwd)/.docker_build
DOCKER_CMD=$(DOCKER_BUILD)/chat

$(DOCKER_CMD): clean
	mkdir -p $(DOCKER_BUILD)
	$(GO_BUILD_ENV) go build -v -o $(DOCKER_CMD) .

clean:
	rm -rf $(DOCKER_BUILD)

heroku: $(DOCKER_CMD)
	heroku container:push web

run-chat-dev:
	export ENV='development' && export PORT='8080' && make build && ./main

fmt:
	gofmt -l -s -w *.go && gofmt  -l -s -w data/. && gofmt -l -s -w  handlers/. && gofmt -l -s -w  core/. && gofmt -l -s -w  authentication/.

lint:
	golint . && golint data/. && golint handlers/. && golint core/. && golint authentication/.

build:
	go build main.go hub.go conn.go
