.PHONY: build dev clean templ css all run

# Build everything
all: templ css build

# Generate templ templates
templ:
	@echo "Generating templ templates..."
	@cd src && templ generate

# Build Tailwind CSS
css:
	@echo "Building Tailwind CSS..."
	@./tailwindcss -i ./src/static/css/input.css -o ./src/static/css/output.css --minify

# Build Go binary
build:
	@echo "Building Go binary..."
	@go build -o bin/server src/main.go

# Development mode with hot reload
dev:
	@echo "Starting development server with Air..."
	@air

# Run the built server
run:
	@echo "Starting server..."
	@./bin/server

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/ tmp/
	@find src/templates -name "*_templ.go" -delete
	@echo "Clean complete"

# Full rebuild (clean + build)
rebuild: clean all

# Help
help:
	@echo "Available commands:"
	@echo "  make all      - Build everything (templ + css + go)"
	@echo "  make templ    - Generate templ templates only"
	@echo "  make css      - Build Tailwind CSS only"
	@echo "  make build    - Build Go binary only"
	@echo "  make dev      - Start development server with Air"
	@echo "  make run      - Run the built server"
	@echo "  make clean    - Remove build artifacts"
	@echo "  make rebuild  - Clean and rebuild everything"
