.PHONY: lint
lint:
	@golangci-lint run

.PHONY: compile
compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o bin/users ./cmd/ports/main.go
