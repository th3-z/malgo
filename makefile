SRC = $(shell find . -type f -name '*.go')

.PHONY: build clean run fmt

build:
	@go build -o mal-sqlite-migrate mal_sqlite_migrate.go
	@chmod +x ./mal-sqlite-migrate

run:
	@./mal-sqlite-migrate

clean:
	-@rm -f ./mal-sqlite-migrate
	-@rm -f ./*.sqlite

fmt:
	@gofmt -l -w $(SRC)

