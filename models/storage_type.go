package models

import (
	"github.com/th3-z/malgo/storage"
)

type StorageType struct {
	Id int
	Name         string
}

func NewStorageType(db storage.Queryer, storageType string) *StorageType {
	query := `
        INSERT OR IGNORE INTO review_storage_type (
            name
        ) VALUES (
            ?
        )
    `

	storageTypeId := storage.PreparedExec(
		db, query, storageType,
	)
	return GetStorageType(db, storageTypeId)
}

func GetStorageType(db storage.Queryer, storageTypeId int64) *StorageType {
	query := `
        SELECT
            review_storage_type_id, name
         FROM
            review_storage_type
        WHERE
            review_storage_type_id = ?
    `
	row := storage.PreparedQueryRow(
		db, query, storageTypeId,
	)
	var storageType StorageType
	row.Scan(&storageType.Id, &storageType.Name)

	return &storageType
}