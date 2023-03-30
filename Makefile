BINARY_NAME=golexer
build:
	go build -o target/${BINARY_NAME} cmd/golexer/main.go

lint:
	golangci-lint run

test:
	go test ./...
