package models

import (
	"github.com/th3-z/malgo/storage"
)

type ReviewStatus struct {
	Id int
	Name         string
}

func NewUserStatus(db storage.Queryer, reviewStatus string) *ReviewStatus {
	query := `
        INSERT OR IGNORE INTO review_status (
            name
        ) VALUES (
            ?
        )
    `

	reviewStatusId := storage.PreparedExec(
		db, query, reviewStatus,
	)
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