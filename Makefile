GO ?= go
BIN_DIR := dist

.PHONY: all build server readgit catgit test clean

all: build

build: server gitreader readgit catgit

server:
	mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/list-repos .

readgit:
	mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/readgit ./cmd/readgit

catgit:
	mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/catgit ./cmd/catgit

test:
	$(GO) test ./...

clean:
	rm -f $(BIN_DIR)/list-repos $(BIN_DIR)/readgit $(BIN_DIR)/catgit
	rmdir $(BIN_DIR) 2>/dev/null || true
