package storage

import (
    "database/sql"
)

func AddUser(db *sql.DB, name string) int64 {
    query  := `
        INSERT INTO user (name)
        VALUES (?)
    `

    insert_id := PreparedQuery(db, query, name)
    return insert_id
}

