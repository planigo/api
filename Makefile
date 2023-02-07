start:
	go run main.go

dev:
	air

install-deps:
	brew install golangci-lint

lint:
	golangci-lint run