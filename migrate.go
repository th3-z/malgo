package main

import (
    "database/sql"
	"github.com/mattn/go-sqlite3"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"io"
)

const (
    schemaFile = "schema.sql"
)

func initDB(filepath string) *sql.DB {
    db, err := sql.Open("sqlite3", filepath)

    if err != nil {
        panic(err)
    }

    return db
}

func createSchema(db *sql.DB) {
    
}

func migrate(db *sql.DB) {
    createSchema(db)

    _, err := db.Exec(query)

    if err != nil {
        panic(err)
    }
}

func main() {
    // TODO: output filename parameter
	db = initDB("my-anime-list.sqlite")
	migrate(db)

    print("Done.")
}
