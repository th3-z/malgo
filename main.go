package main

import (
	"./path"
	"./storage"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
)

func migrate(db *sql.DB, xml string) {
	user := path.GetUser(xml)
	user_id := storage.AddUser(db, user.UserName)

	animeList := path.GetAnimeList(xml)

	for _, anime := range animeList {
		fmt.Printf("SeriesTitle: %s - %d (%s)\n", anime.SeriesTitle, anime.MyScore, anime.MyTags)
	}

	fmt.Printf("sqliteId: %d\n", user_id)
}

func main() {
	var outputFilename string
	flag.StringVar(&outputFilename, "o", "output.sqlite", "Output file")

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		panic("No input file specified")
	}

	inputFilename := args[0]
	xmlBytes, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		panic(err)
	}
	xml := string(xmlBytes)

	db := storage.InitDB("output.sqlite")
	defer db.Close()
	storage.CreateSchema(db)
	migrate(db, xml)

	fmt.Printf("\nMigration complete!\n")
}
