package models

import (
	"github.com/th3-z/malgo/storage"
)

type StorageType struct {
	Id int
	Name         string
}

func NewStorageType(db storage.Queryer, name string) *StorageType {
	query := `
        INSERT INTO review_storage_type (
            name
        ) VALUES (
            ?
        )
    `

	storageTypeId, err := storage.PreparedExec(
		db, query, name,
	)
	if err != nil {
		return SearchStorageType(db, name)
	}
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

func SearchStorageType(db storage.Queryer, name string) *StorageType {
	query := `
        SELECT
            review_storage_type_id, name
         FROM
            review_storage_type
        WHERE
            name = ?
    `
	row := storage.PreparedQueryRow(
		db, query, name,
	)
	var storageType StorageType
	row.Scan(&storageType.Id, &storageType.Name)

	return &storageType
}