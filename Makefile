start:
	go run main.go

dev:
	air

deps:
	go install github.com/cosmtrek/air@latest && \
  	go install gotest.tools/gotestsum

test:
	gotestsum --format short-verbose -- -v ./tests
