build:
	@go build -o bin/api

# build and run..
run: build
	@./bin/api

test:
	@go test -v ./...