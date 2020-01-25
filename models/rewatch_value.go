package models

import (
	"github.com/th3-z/malgo/storage"
)

type RewatchValue struct {
	Id int
	Name         string
}

func NewRewatchValue(db storage.Queryer, rewatchValue string) *RewatchValue {
	query := `
        INSERT OR IGNORE INTO review_rewatch_value (
            name
        ) VALUES (
            ?
        )
    `

	rewatchValueId := storage.PreparedExec(
		db, query, rewatchValue,
	)
	return GetRewatchValue(db, rewatchValueId)
}

func GetRewatchValue(db storage.Queryer, rewatchValueId int64) *RewatchValue {
	query := `
        SELECT
            review_rewatch_value_id, name
         FROM
            review_rewatch_value
        WHERE
            review_rewatch_value_id = ?
    `
	row := storage.PreparedQueryRow(
		db, query, rewatchValueId,
	)
	var rewatchValue RewatchValue
	row.Scan(&rewatchValue.Id, &rewatchValue.Name)

	return &rewatchValue
}