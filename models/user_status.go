package models

import(
    "github.com/th3-z/malgo/storage"
)

type UserStatus struct {
    UserStatusId int
    Name string
}

func AddUserStatus(db storage.Queryer, userStatus string) int64 {
    var userStatusId int64
    query := `
        SELECT
            name
         FROM
            user_status
        WHERE
            name = ?
    `
    err := storage.PreparedQueryRow(
        db, query,
        userStatus,
    ).Scan(&userStatusId)

    if err != nil {
        query = `
            INSERT OR IGNORE INTO user_status (
                name
            ) VALUES (
                ?
            )
        `

        userStatusId = storage.PreparedExec(
            db, query,
            userStatus,
        )
    }

    return userStatusId
}
