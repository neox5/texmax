BIN_DIR=bin

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/texmax ./cmd/texmax-cli

clean:
	rm -rf $(BIN_DIR)

.PHONY: build clean
