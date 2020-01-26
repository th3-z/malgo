package models

import (
	"github.com/th3-z/malgo/storage"
	"time"
)

type Review struct {
	Id              int64
	Series          *Series
	WatchedEpisodes int
	StartDate       time.Time
	FinishDate      time.Time
	Rated           int
	Score           int
	Dvd             int
	Storage         *StorageType
	Status          *ReviewStatus
	Comments        string
	TimesWatched    int
	RewatchValue    *RewatchValue
	Tags            string
	Rewatching      int
	RewatchingEp    int
}

func NewReview(db storage.Queryer, userId int64, seriesId int64) *Review {
	query := `
        INSERT INTO review(
            user_id,
            series_id
        ) VALUES (?, ?)
    `

	_, err := storage.PreparedExec(
		db, query, userId, seriesId,
	)
	if err != nil {
		return GetReview(db, userId, seriesId)
	}
	return GetReview(db, userId, seriesId)
}

func GetReview(db storage.Queryer, userId int64, seriesId int64) *Review {
	query := `
        SELECT
            review_id,
            watched_episodes,
            start_date,
            finish_date,
            rated,
            score,
            dvd,
            review_storage_type_id,
            review_status_id,
            comments,
            times_watched,
            review_rewatch_value_id,
            tags,
            rewatching,
            rewatching_ep
        FROM
            review
        WHERE
            user_id = ?
            AND series_id = ?
    `

	row := storage.PreparedQueryRow(
		db, query, userId, seriesId,
	)
	var review Review
	var storageTypeId int64
	var userStatusId int64
	var rewatchValueId int64

	var startDate int64
	var finishDate int64

	row.Scan(
		&review.Id, &review.WatchedEpisodes, &startDate,
		&finishDate, &review.Rated, &review.Score, &review.Dvd,
		&storageTypeId, &userStatusId, &review.Comments, &review.TimesWatched,
		&rewatchValueId, &review.Tags, &review.Rewatching, &review.RewatchingEp,
	)

	if startDate != 0 {
		review.StartDate = time.Unix(startDate, 0)
	}
	if finishDate != 0 {
		review.FinishDate = time.Unix(finishDate, 0)
	}

	if seriesId != 0 {
		review.Series = GetSeries(db, seriesId)
	}
	if storageTypeId != 0 {
		review.Storage = GetStorageType(db, storageTypeId)
	}
	if userStatusId != 0 {
		review.Status = GetUserStatus(db, userStatusId)
	}
	if rewatchValueId != 0 {
		review.RewatchValue = GetRewatchValue(db, rewatchValueId)
	}

	return &review
}

func getUserReviews(db storage.Queryer, userId int64) []*Review {
	query := `
		SELECT
			series_id
		FROM
			review
		WHERE
			user_id = ?
	`
	rows := storage.PreparedQuery(db, query, userId)
	var reviews []*Review
	for rows.Next() {
		var seriesId int64
		err := rows.Scan(&seriesId)
		if err != nil {
			panic(err)
		}
		reviews = append(reviews, GetReview(db, userId, seriesId))
	}

	return reviews
}

func (review *Review) Update(db storage.Queryer) {
	query := `
        UPDATE review SET
            series_id = ?,
            review_storage_type_id = ?,
            review_status_id = ?,
			review_rewatch_value_id = ?,

            watched_episodes = ?,
            start_date = ?,
            finish_date = ?,
            rated = ?,
            score = ?,
            dvd = ?,
            comments = ?,
            tags = ?,
            times_watched = ?,
            rewatching = ?,
            rewatching_ep = ?
        WHERE
			review_id = ?
    `

	_, err := storage.PreparedExec(
		db, query,
		review.Series.Id,
		review.Storage.Id,
		review.Status.Id,
		review.RewatchValue.Id,

		review.WatchedEpisodes,
		review.StartDate.Unix(),
		review.FinishDate.Unix(),
		review.Rated,
		review.Score,
		review.Dvd,
		review.Comments,
		review.Tags,
		review.TimesWatched,
		review.Rewatching,
		review.RewatchingEp,
		review.Id,
	)
	if err != nil {
		panic(err)
	}

	review.Storage.Update(db)
	review.Status.Update(db)
	review.RewatchValue.Update(db)
	review.Series.Update(db)
}
