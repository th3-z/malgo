package models

import(
    "fmt"
    "mal-sqlite-migrate/storage"
)

type Anime struct {
	SeriesTitle       string
	SeriesAnimeDbId   int
	SeriesType        string
	SeriesEpisodes    int
	MyWatchedEpisodes int
	MyStartDate       int
	MyFinishDate      int
	MyScore           int
	MyStorage         string
	MyStatus          string
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
    print("w")
    seriesTypeId := AddSeriesType(db, anime.SeriesType)

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

    fmt.Printf("Type - %s, TypeId- %d\n", anime.SeriesType, seriesTypeId)

	return storage.PreparedExec(
        db, query,
        anime.SeriesTitle,
        anime.SeriesAnimeDbId,
        anime.SeriesEpisodes,
        seriesTypeId,
    )
}

func AddSeriesType(db storage.Queryer, seriesType string) int64 {
    var seriesId int64
    query := `
        SELECT
            series_type_id
         FROM
            series_type
        WHERE
            name = ?
    `
    err := storage.PreparedQueryRow(
        db, query,
        seriesType,
    ).Scan(&seriesId)

    if err != nil {
        query = `
            INSERT OR IGNORE INTO series_type (
                name
            ) VALUES (
                ?
            )
        `

        seriesId = storage.PreparedExec(
            db, query,
            seriesType,
        )
    }

    return seriesId
}

func SetSeriesType(db storage.Queryer, seriesId int64, seriesTypeId int64) int64 {
    query := `
        UPDATE series SET
            series_type_id = ?
        WHERE
            series_id = ?
    `

    return storage.PreparedExec(
        db, query,
        seriesTypeId,
        seriesId,
    )
}

func SetSeriesStatus(db storage.Queryer, seriesId int64, seriesStatusId int64) int64 {
    query := `
        UPDATE series SET
            series_status_id = ?
        WHERE
            series_id = ?
    `

    return storage.PreparedExec(
        db, query,
        seriesStatusId,
        seriesId,
    )
}
