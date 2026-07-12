BINARY := dataset-factory

.PHONY: build run fmt vet test race lint check tidy clean

build:
	go build -o bin/$(BINARY) ./cmd/dataset-factory

run:
	go run ./cmd/dataset-factory

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...

race:
	go test -race ./...

lint:
	golangci-lint run

tidy:
	go mod tidy

clean:
	rm -rf bin

build:
	go build ./cmd/dataset-factory

check: fmt vet test race build