# Build the application and install any dependencies if necessary
CC=go
BUILDTAGS=-o
DIR=.
EXECUTABLE=sanalyzer
DEP=github.com/fatih/color

build:
	@echo "+ $@"
	@echo "+ Downloading dependencies. Please wait..."
	@go get -u $(DEP)
	@echo "+ Building the project..."
	@$(CC) build $(BUILDTAGS) $(EXECUTABLE) $(DIR)
	@echo "+ Build complete!\n"
	@echo "+ Run by typing: ./"$(EXECUTABLE)

clean:
	@echo "+ $@"
	@rm -rf sanalyzer
	@echo "Cleanup complete!"

all: build clean
