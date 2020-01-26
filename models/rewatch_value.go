package models

import (
	"github.com/th3-z/malgo/storage"
)

type RewatchValue struct {
	Id int
	Name         string
}

func NewRewatchValue(db storage.Queryer, name string) *RewatchValue {
	query := `
        INSERT INTO review_rewatch_value (
            name
        ) VALUES (
            ?
        )
    `

	rewatchValueId, err := storage.PreparedExec(
		db, query, name,
	)
	if err != nil {
		return SearchRewatchValue(db, name)
	}
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

func SearchRewatchValue(db storage.Queryer, name string) *RewatchValue {
	query := `
        SELECT
            review_rewatch_value_id, name
         FROM
            review_rewatch_value
        WHERE
            name = ?
    `
	row := storage.PreparedQueryRow(
		db, query, name,
	)
	var rewatchValue RewatchValue
	row.Scan(&rewatchValue.Id, &rewatchValue.Name)

	return &rewatchValue
}

func (rewatchValue *RewatchValue) Update (db storage.Queryer) {
	query := `
        UPDATE review_rewatch_value SET
			name = ?
        WHERE
            review_rewatch_value_id  = ?
    `
	_, err := storage.PreparedExec(
		db, query, rewatchValue.Name, rewatchValue.Id,
	)
	if err != nil {
		panic(err)
	}
}
