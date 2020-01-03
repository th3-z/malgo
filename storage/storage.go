package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)


type Queryer interface {
    Query(string, ...interface{}) (*sql.Rows, error)
    QueryRow(string, ...interface{}) *sql.Row
    Prepare(string) (*sql.Stmt, error)
    Exec(string, ...interface{}) (sql.Result, error)
}


func InitDB(filepath string) *sql.DB {
	if _, err := os.Stat(filepath); err == nil {
		os.Remove(filepath)
	}
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	return db
}

func CreateSchema(db Queryer) {
	query := schema // storage/schema.go
	_, err := db.Exec(query)

	if err != nil {
		panic(err)
	}
}

func PreparedQuery(db Queryer, query string, args ...interface{}) int64 {
	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(args...)
	if err != nil {
		panic(err)
	}

	insertId, err := res.LastInsertId()
	if err == nil {
		return insertId
	}

	affectedRows, err := res.RowsAffected()
	if err == nil {
		return affectedRows
	}

	return 0
}
