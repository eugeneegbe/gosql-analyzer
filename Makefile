# Build the application and install any dependencies if necessary
CC=go
BUILDTAGS=-o

all: build clean

build:
	@echo "+ $@"
	@$(CC) build $(BUILDTAGS) sanalyzer .
	@echo "Build complete!"

clean:
	@echo "+ $@"
	@rm -rf sanalyzer
	@echo "Cleanup complete!"