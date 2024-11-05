.PHONY: run test format 

run:
	@go run cmd/main.go

test:
	@go test -v ./...

format:
	@go fmt ./...
