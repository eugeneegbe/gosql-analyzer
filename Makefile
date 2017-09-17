# Build the application

all: clean build

build:
	@echo "+ $@"
	@go build -o sanalyzer .
	@echo "Build complete!"

clean:
	@echo "+ $@"
	@rm -rf sanalyzer
	@echo "Cleanup complete!"