start:
	go run main.go

dev:
	air

install-deps:
	go install github.com/cosmtrek/air@latest && \
	brew install golangci-lint

lint:
	golangci-lint run

config:
	git config core.hooksPath .githooks

deps:
	go install github.com/cosmtrek/air@latest && \
  	go install gotest.tools/gotestsum

test:
	gotestsum --format testname ./tests
