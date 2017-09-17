# Build the application

all: clean build install

build:
	@echo "+ $@"
	@go build .

clean:
	@echo "+ $@"
	@rm -rf sanalyze

install:
	@echo "+ $@"
	@go install .
