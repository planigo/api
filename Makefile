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