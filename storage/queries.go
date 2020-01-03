package storage

import (
	"../path"
	"database/sql"
)

func AddUser(db *sql.DB, user *path.User) int64 {
	query := `
        INSERT INTO user (name)
        VALUES (?)
    `

	insert_id := PreparedQuery(db, query, user.UserName)
	return insert_id
}

func AddAnime(db *sql.DB, anime *path.Anime) int64 {
	query := `
        INSERT INTO series (name)
        VALUES (?)
    `

	insert_id := PreparedQuery(db, query, anime.SeriesTitle)
	return insert_id
}
