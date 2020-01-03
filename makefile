SRC = $(shell find . -type f -name '*.go')

.PHONY: build clean run fmt

build:
	@go build -o mal-sqlite-migrate main.go
	@chmod +x ./mal-sqlite-migrate

clean:
	-@rm -f ./mal-sqlite-migrate
	-@rm -f ./*.sqlite

fmt:
	@gofmt -l -w $(SRC)

