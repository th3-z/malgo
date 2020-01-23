package malgo

import (
    "github.com/th3-z/malgo/utils"
    "github.com/th3-z/malgo/parsers"
    "database/sql"
)

func MigrateFile(db *sql.DB, filename string) {
    MigrateString(db, utils.FileToStr(filename))
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

