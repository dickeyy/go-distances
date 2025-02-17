#!/bin/bash

# Directory containing the test data files
TEST_DATA_DIR="./test-data"

# Path to the Go program executable
GO_PROGRAM="./go-distances" # Replace with the actual path to your Go program (compiled executable)

# Check if the test data directory exists
if [ ! -d "$TEST_DATA_DIR" ]; then
  echo "Error: Test data directory '$TEST_DATA_DIR' not found."
  exit 1
fi

# Check if the Go program exists
if [ ! -x "$GO_PROGRAM" ]; then
  echo "Error: Go program '$GO_PROGRAM' not found or is not executable."
  exit 1
fi

# Loop through each file in the test data directory
for file in "$TEST_DATA_DIR"/*.json; do  # Only process .json files
  # Check if it's a regular file
  if [ -f "$file" ]; then
    echo "Processing file: $file"

    # Provide "y" and the file path to the Go program via stdin
    printf "y\n%s\n" "$file" | "$GO_PROGRAM"

    echo "Finished processing: $file"
  fi
done

echo "Finished processing all files."
