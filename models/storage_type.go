package models

import (
	"github.com/th3-z/malgo/storage"
)

type StorageType struct {
	Id   int
	Name string
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

func (storageType *StorageType) Update(db storage.Queryer) {
	query := `
        UPDATE review_storage_type SET
			name = ?
        WHERE
            review_storage_type_id = ?
    `
	_, err := storage.PreparedExec(
		db, query, storageType.Name, storageType.Id,
	)
	if err != nil {
		panic(err)
	}
}

func (storageType *StorageType) Delete(db storage.Queryer) {
	query := `
        DELETE FROM review_storage_type
        WHERE
            review_storage_type_id = ?
    `
	_, err := storage.PreparedExec(
		db, query, storageType.Id,
	)
	if err != nil {
		panic(err)
	}
}
