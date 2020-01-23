package models

import (
	"github.com/th3-z/malgo/storage"
)

type Anime struct {
	SeriesId          int
	SeriesTitle       string
	SeriesAnimeDbId   int
	SeriesType        SeriesType
	SeriesEpisodes    int
	MyWatchedEpisodes int
	MyStartDate       int
	MyFinishDate      int
	MyScore           int
	MyStorage         int // TODO: Storage model
	MyStatus          UserStatus
	MyComments        string
	MyTimesWatched    int
	MyRewatchValue    int
	MyTags            string
	MyRewatching      int
	MyRewatchingEp    int

	MyId           int    // Unused
	MyDvd          string // Unused
	MyRated        int    // Unused
	UpdateOnImport int    // Unused
}

func AddSeries(db storage.Queryer, anime *Anime) int64 {
	seriesTypeId := AddSeriesType(db, anime.SeriesType.SeriesTypeId)

	query := `
        INSERT INTO series (
            name,
            animedb_id,
            episodes,
            series_type_id
        ) VALUES (
            ?,
            ?,
            ?,
            ?
        )
    `

	return storage.PreparedExec(
		db, query,
		anime.SeriesTitle,
		anime.SeriesAnimeDbId,
		anime.SeriesEpisodes,
		seriesTypeId,
	)
}

func GetAnimeList(db storage.Queryer) []Anime {
	var animeList []Anime

	query := `
        SELECT 
            name,
            animedb_id,
            episodes
        FROM series
    `
	rows := storage.PreparedQuery(
		db, query,
	)
	defer rows.Close()

	for rows.Next() {
		var anime Anime
		err := rows.Scan(
			&anime.SeriesTitle, &anime.SeriesAnimeDbId, &anime.SeriesEpisodes,
		)
		if err != nil {
			panic(err)
		}

		animeList = append(animeList, anime)
	}

	return animeList
}

func SetSeriesType(db storage.Queryer, anime *Anime, seriesTypeId int64) int64 {
	query := `
        UPDATE series SET
            series_type_id = ?
        WHERE
            series_id = ?
    `

	return storage.PreparedExec(
		db, query,
		seriesTypeId,
		anime.SeriesId,
	)
}
