.PHONY: all build clean run

all: clean build run

build:
	@go build -o bin/sn-update-set-extractor cmd/sn-update-set-extractor/*.go

clean:
	@rm -rf bin dist

run:
	@bin/sn-update-set-extractor
