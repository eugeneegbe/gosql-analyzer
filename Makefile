# Build the application and install any dependencies if necessary
CC=go
BUILDTAGS=-o
DIR=.

build:
	@echo "+ $@"
	@$(CC) build $(BUILDTAGS) sanalyzer $(DIR)
	@echo "Build complete!"

clean:
	@echo "+ $@"
	@rm -rf sanalyzer
	@echo "Cleanup complete!"

all: build clean
