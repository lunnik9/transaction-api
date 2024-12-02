.DEFAULT_GOAL := build_and_run

build:
	go mod tidy && go mod download && \
	go build -o ./bin/transaction-parser-api/${BINARY_NAME} ./cmd/transaction-parser-api/main.go

run:
	./bin/transaction-parser-api/${BINARY_NAME}

build_and_run: build run

test:
	go test ./...

lint:
	golangci-lint run

clean:
	go clean
	go clean --testcache
	rm ./bin/transaction-parser-api/*