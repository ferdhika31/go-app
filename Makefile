.PHONY : run format clean install build

run:
	go run ./internal/main.go

watch:
	air -c .air.toml

format:
	gofmt -s -w .

clean:
	go mod tidy

install:
	go mod download

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags musl -o bin/main ./internal/main.go

start:
	./bin/main

test:
	go test -coverprofile=coverages/coverage.out ./tests/...
	go tool cover -html=coverages/coverage.out -o coverages/coverage.html

