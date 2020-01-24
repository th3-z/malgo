package models

import (
	"github.com/th3-z/malgo/storage"
)

type Series struct {
	Id          int64
	Title       string
	AnimedbId   int
	Type        *SeriesType
	Episodes    int
}

func NewSeries(db storage.Queryer, seriesTitle string) *Series {
	query := `
        INSERT INTO series (
            name
        ) VALUES (
            ?
        )
    `

    seriesId := storage.PreparedExec(
		db, query, seriesTitle,
	)

    return GetSeries(db, seriesId)
}

func GetSeries(db storage.Queryer, seriesId int64) *Series {
    query := `
        SELECT
            series_id,
            name,
            animedb_id,
            episodes,
            series_type_id
        FROM
            series
        WHERE
            series_id = ?
    `

    row := storage.PreparedQueryRow(
        db, query, seriesId,
    )
    var series Series
    var seriesTypeId int64
    row.Scan(
        &series.Id, &series.Title, &series.AnimedbId, &series.Episodes,
        &seriesTypeId,
    )

    if seriesTypeId != 0 {
        series.Type = GetSeriesType(db, seriesTypeId)
        return &series
    }

    return &series
}

