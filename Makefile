.PHONY: build run clean test

BINARY=bin/camera-app

build:
	go build -o $(BINARY) ./cmd/camera-app

run: build
	./$(BINARY) -driver dummy -watermark testdata/logo.png

test:
	go test -v ./...

clean:
	rm -rf $(BINARY) output/