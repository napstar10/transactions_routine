#!/bin/bash

# Define project directory and binary name
PROJECT_DIR=$(pwd)
BINARY_NAME="transactions_routine"

# Run the tests first
echo "Running tests..."
go test ./... -v

# Capture the test result
TEST_RESULT=$?

if [ $TEST_RESULT -ne 0 ]; then
  echo "Tests failed. Exiting."
  exit 1
fi

# Build the project if tests pass
echo "Tests passed. Building the project..."

# Change to the main directory before building the app

# Build the Go project in the main directory
go build -o $BINARY_NAME ./main

if [ $? -ne 0 ]; then
  echo "Build failed. Exiting."
  exit 1
fi

# Run the project (start the server)
echo "Starting the server..."
./$BINARY_NAME &


echo "Server is running. Press Ctrl+C to stop."