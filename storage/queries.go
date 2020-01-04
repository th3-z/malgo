package storage

import (
    "fmt"
	"../path"
)


func AddUser(db Queryer, user *path.User) int64 {
	query := `
        INSERT INTO user (name)
        VALUES (?)
    `

	insert_id := PreparedExec(db, query, user.UserName)
	return insert_id
}



func AddSeriesType(db Queryer, seriesType string) int64 {
    var seriesId int64
    query := `
        SELECT
            series_type_id
         FROM
            series_type
        WHERE
            name = ?
    `
    err := PreparedQueryRow(
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

        seriesId = PreparedExec(
            db, query,
            seriesType,
        )
    }

    return seriesId
}

func AddSeries(db Queryer, anime *path.Anime) int64 {
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

	return PreparedExec(
        db, query,
        anime.SeriesTitle,
        anime.SeriesAnimeDbId,
        anime.SeriesEpisodes,
        seriesTypeId,
    )
}

func SetSeriesType(db Queryer, seriesId int64, seriesTypeId int64) int64 {
    query := `
        UPDATE series SET
            series_type_id = ?
        WHERE
            series_id = ?
    `

    return PreparedExec(
        db, query,
        seriesTypeId,
        seriesId,
    )
}

func SetSeriesStatus(db Queryer, seriesId int64, seriesStatusId int64) int64 {
    query := `
        UPDATE series SET
            series_status_id = ?
        WHERE
            series_id = ?
    `

    return PreparedExec(
        db, query,
        seriesStatusId,
        seriesId,
    )
}

func AddUserAnime(db Queryer, userId int64, seriesId int64) int64 {
	query := `
        INSERT INTO user_series (
            user_id,
            series_id
        ) VALUES (
            ?,
            ?
        )
    `

	return PreparedExec(db, query, userId, seriesId)
}

