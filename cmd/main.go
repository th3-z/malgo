package main

import (
	"github.com/th3-z/malgo"
	"github.com/th3-z/malgo/storage"
	"flag"
	"fmt"
)

func main() {
	var outputFilename string
	flag.StringVar(&outputFilename, "o", "output.sqlite", "Output file")

	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		panic("No input file specified")
	}

	inputFilename := args[0]

	db := storage.InitDB(outputFilename)
	defer db.Close()

	storage.CreateSchema(db)
	malgo.MigrateFile(db, inputFilename)

	fmt.Printf("\nMigration complete!\n")
}
