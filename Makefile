BINARY_FILE_PATH=bin/main
BIN_DIR_PATH=bin
MAIN_FILE_PATH=todo-app/main/main.go

default: build

build:
#	@if  ! ./build-up-to-date.sh; then \
  		make fmt; \
  		make test; \
  		echo "Building the binary..."; \
        go build -o $(BINARY_FILE_PATH) $(MAIN_FILE_PATH); \
  	fi;
	echo "Building the binary..."; \
	CGO_ENABLED=0 go build -o $(BINARY_FILE_PATH) $(MAIN_FILE_PATH); \

fmt:
	@echo "Formatting code..."
	go fmt ./...

test:
	@echo "Running tests..."
	go test -v ./...

clean:
	@echo "Cleaning up..."
	go clean
	rm -fR $(BIN_DIR_PATH)
	
deps:
	@echo "Installing dependencies..."
	go get -v ./...

.PHONY: build run fmt test clean deps
