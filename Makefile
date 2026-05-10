run:
	@go run ./cmd/app/main.go

test:
	@go test ./...

test-v:
	@go test -v ./...

test-race:
	@go test -race ./...

test-cover:
	@go test -cover ./...