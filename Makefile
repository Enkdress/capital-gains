build:
	@go build -o bin/capitalgains ./cmd 

run:
	./bin/capitalgains < input.txt

test:
	@go test -v ./...
