package models

import (
	"github.com/th3-z/malgo/storage"
)

type SeriesType struct {
	Id   int
	Name string
}

func NewSeriesType(db storage.Queryer, name string) *SeriesType {
	query := `
        INSERT INTO series_type (
            name
        ) VALUES (
            ?
        )
    `

	seriesTypeId, err := storage.PreparedExec(
		db, query, name,
	)
	if err != nil {
		return SearchSeriesType(db, name)
	}
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

func SearchSeriesType(db storage.Queryer, name string) *SeriesType {
	query := `
        SELECT
            series_type_id, name
         FROM
            series_type
        WHERE
            name = ?
    `
	row := storage.PreparedQueryRow(
		db, query, name,
	)
	var seriesType SeriesType
	row.Scan(&seriesType.Id, &seriesType.Name)

	return &seriesType
}

func (seriesType *SeriesType) Update(db storage.Queryer) {
	query := `
        UPDATE series_type SET
			name = ?
        WHERE
            series_type_id = ?
    `
	_, err := storage.PreparedExec(
		db, query, seriesType.Name, seriesType.Id,
	)
	if err != nil {
		panic(err)
	}
}
