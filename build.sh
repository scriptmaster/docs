#!/bin/bash

# Build script for docs server

echo "Building documentation server..."

# Build for Windows
echo "Building for Windows (amd64)..."
GOOS=windows GOARCH=amd64 go build -o docs.exe
if [ $? -eq 0 ]; then
    echo "✓ Windows build complete: docs.exe"
else
    echo "✗ Windows build failed"
    exit 1
fi

# Build for Linux
echo "Building for Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -o docs-linux
if [ $? -eq 0 ]; then
    echo "✓ Linux build complete: docs-linux"
else
    echo "✗ Linux build failed"
    exit 1
fi

# Build for macOS
echo "Building for macOS (amd64)..."
GOOS=darwin GOARCH=amd64 go build -o docs-macos
if [ $? -eq 0 ]; then
    echo "✓ macOS build complete: docs-macos"
else
    echo "✗ macOS build failed"
    exit 1
fi

echo ""
echo "All builds completed successfully!"
echo "Files created:"
ls -lh docs.exe docs-linux docs-macos 2>/dev/null || true
