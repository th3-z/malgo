package models

import (
	"github.com/th3-z/malgo/storage"
)

type User struct {
	Id               int64
	Name             string
}

func NewUser(db storage.Queryer, name string) *User {
	query := `
        INSERT INTO user (name)
        VALUES (?)
    `

	userId := storage.PreparedExec(db, query, name)
	return GetUser(db, userId)
}

func GetUser(db storage.Queryer, userId int64) *User {
    query := `
        SELECT
            user_id,
            name
        FROM
            user
        WHERE
            user_id = ?
    `

    row := storage.PreparedQueryRow(
        db, query, userId,
    )
    var user User
    row.Scan(
        &user.Id, &user.Name,
    )

    return &user
}

