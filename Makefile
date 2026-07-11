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

check: fmt vet test race