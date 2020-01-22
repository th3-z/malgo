package models

import(
    "mal-sqlite-migrate/storage"
)

type User struct {
	UserId               int
	UserName             string
	UserExportType       int // Unused
	UserTotalAnime       int // Unused
	UserTotalWatching    int // Unused
	UserTotalCompleted   int // Unused
	UserTotalOnHold      int // Unused
	UserTotalDropped     int // Unused
	UserTotalPlanToWatch int // Unused
}

func AddUser(db storage.Queryer, user *User) int64 {
	query := `
        INSERT INTO user (name)
        VALUES (?)
    `

	insert_id := storage.PreparedExec(db, query, user.UserName)
	return insert_id
}



func AddUserAnime(db storage.Queryer, userId int64, seriesId int64) int64 {
	query := `
        INSERT INTO user_series (
            user_id,
            series_id
        ) VALUES (
            ?,
            ?
        )
    `

	return storage.PreparedExec(db, query, userId, seriesId)
}

