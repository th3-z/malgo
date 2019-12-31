package storage

import (
    "database/sql"
	_"github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) *sql.DB {
    db, err := sql.Open("sqlite3", filepath)

    if err != nil {
        panic(err)
    }

    return db
}

func CreateSchema(db *sql.DB) {
    query := schema
    _, err := db.Exec(query)

    if err != nil {
        panic(err)
    }
}

func ins(db *sql.DB, query string, args...string) int64 {
    stmt, err := db.Prepare(query)
    if (err != nil) {
        panic(err)
    }

    res, err := stmt.Exec(args...)
    if (err != nil) {
        panic(err)
    }

    id, _ := res.LastInsertId()
    return id
}


func AddUser(db *sql.DB, name string) {
    query  := `
        INSERT INTO user (name)
        VALUES (?)
    `

    insert_id := ins(db, query, name)
    print(string(insert_id))
}

