BINARY := dataset-factory

.PHONY: build
build:
	go build -o bin/$(BINARY) ./cmd/dataset-factory

.PHONY: run
run:
	go run ./cmd/dataset-factory

.PHONY: test
test:
	go test ./...

.PHONY: fmt
fmt:
	gofmt -w .

.PHONY: vet
vet:
	go vet ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: clean
clean:
	rm -rf bin