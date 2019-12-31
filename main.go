package main

import (
    "database/sql"
	_"github.com/mattn/go-sqlite3"
    "./storage"
)


func migrate(db *sql.DB) {
    storage.AddUser(db, "the_z")
}

func main() {
    // TODO: output filename parameter
	db := storage.InitDB("my-anime-list.sqlite")
    storage.CreateSchema(db)

    migrate(db)
    db.Close()

    print("Done.\n")
}

