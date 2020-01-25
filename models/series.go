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

func NewSeries(db storage.Queryer, name string) *Series {
	query := `
        INSERT INTO series (
            name
        ) VALUES (
            ?
        )
    `

    seriesId, err := storage.PreparedExec(
		db, query, name,
	)

    if err != nil {
    	return SearchSeries(db, name)
	}

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

func SearchSeries(db storage.Queryer, name string) *Series {
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
            name = ?
    `

	row := storage.PreparedQueryRow(
		db, query, name,
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

