build:
	@go build -o bin/capitalgains ./cmd 

run: build
	@./bin/capitalgains

test:
	@go test -v ./...
