package models

import(
    "github.com/th3-z/malgo/storage"
)

type SeriesType struct {
    SeriesTypeId int
    Name string
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

