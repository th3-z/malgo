package models

import (
	"github.com/th3-z/malgo/storage"
)

type SeriesType struct {
	Id int
	Name         string
}

func NewSeriesType(db storage.Queryer, seriesType string) *SeriesType {
    query := `
        INSERT OR IGNORE INTO series_type (
            name
        ) VALUES (
            ?
        )
    `

    seriesTypeId := storage.PreparedExec(
		db, query, seriesType,
	)
    return GetSeriesType(db, seriesTypeId)
}

func GetSeriesType(db storage.Queryer, seriesTypeId int64) *SeriesType {
	query := `
        SELECT
            series_type_id, name
         FROM
            series_type
        WHERE
            series_type_id = ?
    `
	row := storage.PreparedQueryRow(
		db, query, seriesTypeId,
    )
    var seriesType SeriesType
    row.Scan(&seriesType.Id, &seriesType.Name)

    return &seriesType
}

