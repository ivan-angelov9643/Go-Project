BINARY_FILE_PATH=bin/main
MAIN_FILE_PATH=main/main.go

default: build

build: fmt test
	@echo "Building the binary..."
	go build -o $(BINARY_FILE_PATH) $(MAIN_FILE_PATH)

run: build
	@echo "Running the app..."
	./bin/main

fmt:
	@echo "Formatting code..."
	go fmt ./...

test:
	@echo "Running tests..."
	go test -v ./...

clean:
	@echo "Cleaning up..."
	go clean
	rm -f $(BINARY_FILE_PATH)
	
deps:
	@echo "Installing dependencies..."
	go get -v ./...

.PHONY: build run fmt test clean deps
