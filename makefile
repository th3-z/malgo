build:
	@go build -o mal-sqlite-migrate mal_sqlite_migrate.go
	@chmod +x ./mal-sqlite-migrate

run:
	@./mal-sqlite-migrate

clean:
	@rm ./mal-sqlite-migrate
	@rm ./output.sqlite

