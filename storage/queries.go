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
        INSERT INTO series (name)
        VALUES (?)
    `

	insert_id := PreparedQuery(db, query, anime.SeriesTitle)
	return insert_id
}

