package main

import (
    "database/sql"
	_"github.com/mattn/go-sqlite3"
    "./storage"
    "fmt"
)


func migrate(db *sql.DB) {
    user_id := storage.AddUser(db, "the_z")

    fmt.Printf("%d\n", user_id)
}

func main() {
    // TODO: output filename parameter
	db := storage.InitDB("my-anime-list.sqlite")
    storage.CreateSchema(db)

    migrate(db)
    db.Close()

    fmt.Printf("Done.\n")
}

