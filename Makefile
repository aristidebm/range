.PHONY: run test format 

run:
	@go run main.go

test:
	@go test -v ./...

format:
	@go fmt ./...
