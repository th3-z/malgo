package models

import(
    "github.com/th3-z/malgo/storage"
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

func AddUserAnime(db storage.Queryer, user *User, seriesId int64) int64 {
	query := `
        INSERT INTO user_series (
            user_id,
            series_id
        ) VALUES (
            ?,
            ?
        )
    `

	return storage.PreparedExec(db, query, User.UserId, seriesId)
}

func SetUserAnimeUserStatus(db storage.Queryer, user *User, anime *Anime, userStatusId int64) int64 {
    query := `
        UPDATE user_series SET
            user_status_id = ?
        WHERE
            user_series_id = ?
            AND user_id = ?
    `

    return storage.PreparedExec(db, userStatusId, anime.SeriesId, user.UserId)
}

