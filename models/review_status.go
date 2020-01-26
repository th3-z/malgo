package models

import (
	"github.com/th3-z/malgo/storage"
)

type ReviewStatus struct {
	Id int
	Name         string
}

func NewUserStatus(db storage.Queryer, name string) *ReviewStatus {
	query := `
        INSERT INTO review_status (
            name
        ) VALUES (
            ?
        )
    `

	reviewStatusId, err := storage.PreparedExec(
		db, query, name,
	)
	if err != nil {
		return SearchUserStatus(db, name)
	}
	return GetUserStatus(db, reviewStatusId)
}

func GetUserStatus(db storage.Queryer, reviewStatusId int64) *ReviewStatus {
	query := `
        SELECT
            review_status_id, name
         FROM
            review_status
        WHERE
            review_status_id = ?
    `
	row := storage.PreparedQueryRow(
		db, query, reviewStatusId,
	)
	var reviewStatus ReviewStatus
	row.Scan(&reviewStatus.Id, &reviewStatus.Name)

	return &reviewStatus
}

func SearchUserStatus(db storage.Queryer, name string) *ReviewStatus {
	query := `
        SELECT
            review_status_id, name
         FROM
            review_status
        WHERE
            name = ?
    `
	row := storage.PreparedQueryRow(
		db, query, name,
	)
	var reviewStatus ReviewStatus
	row.Scan(&reviewStatus.Id, &reviewStatus.Name)

	return &reviewStatus
}

func (reviewStatus *ReviewStatus) Update (db storage.Queryer) {
	query := `
        UPDATE review_status SET
			name = ?
        WHERE
            review_status_id = ?
    `
	_, err := storage.PreparedExec(
		db, query, reviewStatus.Name, reviewStatus.Id,
	)
	if err != nil {
		panic(err)
	}
}
