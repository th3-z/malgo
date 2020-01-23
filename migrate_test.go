package malgo

import (
	"github.com/th3-z/malgo/storage"
    "testing"
)

func TestMigrateFile(t *testing.T) {
	inputFilename := "sample.xml"
	outputFilename := "output.sqlite"

	db := storage.InitDB(outputFilename)
	defer db.Close()

	storage.CreateSchema(db)
	MigrateFile(db, inputFilename)
}
