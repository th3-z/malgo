package main

import (
    "database/sql"
	_"github.com/mattn/go-sqlite3"
    "./storage"
    "./path"
    "fmt"
    "io/ioutil"
)


func migrate(db *sql.DB, xml string) {
    user := path.GetUser(xml)

    user_id := storage.AddUser(db, user.Name)

    fmt.Printf("sqliteId: %d\n", user_id)
    fmt.Printf("username: %s\n", user.Name)
    fmt.Printf("malId: %d\n", user.MalUserId)
}


func main() {
    // TODO: input and output filename parameter
    xmlFilename := "input.xml"
    xmlBytes, err := ioutil.ReadFile(xmlFilename)
    if err != nil {
        panic(err)
    }
    xml := string(xmlBytes)

	db := storage.InitDB("output.sqlite")
    defer db.Close()
    storage.CreateSchema(db)



    migrate(db, xml)

    fmt.Printf("Done.\n")
}

