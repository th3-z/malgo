package storage

import (
	"../path"
)


func AddUser(db Queryer, user *path.User) int64 {
	query := `
        INSERT INTO user (name)
        VALUES (?)
    `

	insert_id := PreparedQuery(db, query, user.UserName)
	return insert_id
}

func AddAnime(db Queryer, anime *path.Anime) int64 {
	query := `
        INSERT INTO series (
            name,
            animedb_id,
            series_type_id,
            series_status_id
        ) VALUES (
            ?,
            ?,
            1,
            1
        )
    `

	insert_id := PreparedQuery(db, query, anime.SeriesTitle, anime.SeriesAnimeDbId)
	return insert_id
}

func AddUserAnime(db Queryer, userId int, seriesId int) int64 {
	query := `
        INSERT INTO user_series (
            user_id,
            series_id
        ) VALUES (
            ?,
            ?
        )
    `

	insert_id := PreparedQuery(db, query, userId, seriesId)
	return insert_id
}
