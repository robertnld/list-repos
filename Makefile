GO ?= go
BIN_DIR := dist

.PHONY: all build server gitreader test clean catgit

all: build

build: server gitreader catgit

server:
	mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/list-repos .

gitreader:
	mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/gitreader ./cmd/gitreader

catgit:
	mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/catgit ./cmd/catgit

test:
	$(GO) test ./...

clean:
	rm -f $(BIN_DIR)/list-repos $(BIN_DIR)/gitreader
	rmdir $(BIN_DIR) 2>/dev/null || true
