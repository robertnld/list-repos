GO ?= go
BIN_DIR := dist

.PHONY: all build server gitreader test clean

all: build

build: server gitreader

server:
	mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/list-repos .

gitreader:
	mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/gitreader ./cmd/gitreader

test:
	$(GO) test ./...

clean:
	rm -f $(BIN_DIR)/list-repos $(BIN_DIR)/gitreader
	rmdir $(BIN_DIR) 2>/dev/null || true
