set dotenv-load

frontend_dir := "frontend"

binary := "bin/server"

# List available recipes
default:
    just --list

# Start backend and frontend
dev:
    #!/bin/bash
    trap 'kill 0' SIGINT
    go run ./cmd/server &
    npm --prefix {{frontend_dir}} run dev &
    wait

# Run Go server only
backend:
    go run ./cmd/server

# Run Svelte dev server only
frontend:
    npm --prefix {{frontend_dir}} run dev

# Install frontend dependencies
install:
    npm --prefix {{frontend_dir}} install

# Production build
build:
    npm --prefix {{frontend_dir}} run build
    go build -o {{binary}} ./cmd/server

# Remove build artifacts
clean:
    rm -rf {{binary}} {{frontend_dir}}/dist
