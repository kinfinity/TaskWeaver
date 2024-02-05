# Makefile TaskWeaver Project

.PHONY: help build clean

help:
	@echo "Available make targets:"
	@echo "  build       - Build TaskWeaver Binaries"
	@echo "  clean       - Clean up generated files"

build:
	go build -o bin/taskweaver cmd/taskweaver/taskweaver.go
	go build -o bin/taskagent cmd/taskagent/taskagent.go

clean:
	@echo "Cleaning up generated files..."
	@# Execute cleanup functions here
	-rm -rf ./bin/

