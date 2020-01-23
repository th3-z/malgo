package malgo

import (
    "github.com/th3-z/malgo/parsers"
    "database/sql"
    "io/ioutil"
)

func MigrateFile(db *sql.DB, filename string) {
    xmlBytes, err := ioutil.ReadFile(filename)
    if err != nil {
        panic(err)
    }

    MigrateString(db, string(xmlBytes))
}

func MigrateString(db *sql.DB, xml string) {
    // Inserts spawn transactions, they are slow, run this in one transaction
    tx, err := db.Begin()
    if (err != nil) {
        panic(err)
    }
    defer tx.Commit()

	malXml := parsers.ParseMalXml(xml)

	for _, anime := range *malXml.AnimeXml {
        print(anime.SeriesTitle)
        print("\n")
		// models.AddSeries(tx, &anime)
        // models.AddUserAnime(tx, &user, anime.SeriesId)
        // statusId := models.AddUserStatus(tx, &user, anime.SeriesStatus)
        // models.SetAnimeUserStatus(tx, &user, &anime, statusId)
	}
}

