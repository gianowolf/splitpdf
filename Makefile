# Variables
BINARY_NAME = splitpdf
TEST_DIR = test
INPUT_PDF = $(TEST_DIR)/input.pdf
CHAPTERS_CSV = $(TEST_DIR)/chapters.csv
PREFIX = Demo

# Default target
all: build

# Build the binary
build:
	go build -o $(BINARY_NAME) .

# Run the tool with test files
run: build
	./$(BINARY_NAME) -i $(INPUT_PDF) -f $(CHAPTERS_CSV) -n $(PREFIX)

# Run with auto-detect mode (only one PDF and CSV in test/)
run-dir: build
	./$(BINARY_NAME) $(TEST_DIR)

# Clean build artifacts
clean:
	rm -f $(BINARY_NAME)

# Remove generated PDFs from test/
clean-test:
	rm -f $(TEST_DIR)/*.pdf
	rm -f $(TEST_DIR)/*.txt

# Reset everything
reset: clean clean-test

.PHONY: all build run run-dir clean clean-test reset
