BINARY=bin/photostamp-cli

.PHONY: build run run-gocv test bench clean

build:
	go build -o $(BINARY) ./cmd

run: build
	./$(BINARY) -driver dummy -watermark testdata/logo.png -margin 10 -scale 0.5

run-gocv: build
	./$(BINARY) -driver gocv -watermark testdata/logo.png -margin 20 -scale 0.5

test:
	go test -v ./...

bench:
	go test -bench=. -run=^$ ./internal/watermark

clean:
	rm -rf bin/ output/