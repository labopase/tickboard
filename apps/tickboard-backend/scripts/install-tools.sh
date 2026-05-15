#!/bin/bash

set -e

echo "Installing tools..."

# sqlc
# https://docs.sqlc.dev/en/stable/overview/install.html
if command -v sqlc &> /dev/null; then
    echo "sqlc already installed"
else
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.30.0
fi

# mockery
# https://docs.mockery.dev/en/stable/installation.html
if command -v mockery &> /dev/null; then
    echo "mockery already installed"
else
    go install github.com/vektra/mockery/v2@latest
fi

echo "All tools installed successfully!"