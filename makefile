.PHONY: all build concat clean

BIN_DIR=bin

all: build concat

build:
	@echo "Building texmax..."
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/texmax ./cmd/texmax-cli

concat:
	@echo "Concating project..."
	@touch ./output.txt
	@./concat_project.sh . output.txt

clean:
	rm -rf $(BIN_DIR)

