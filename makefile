SRC = $(shell find . -type f -name '*.go')

.PHONY: build clean fmt

build:
        @go build -o bin/malgo cmd/main.go
        @chmod +x bin/malgo

clean:
        -@rm -rf ./bin
        -@rm -f ./malgo
        -@rm -f ./*.sqlite

fmt:
        @gofmt -l -w $(SRC)
