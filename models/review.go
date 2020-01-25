package models

import (
	"github.com/th3-z/malgo/storage"
	"time"
)

type Review struct {
	Id               int64
	Series           *Series
    WatchedEpisodes int
    StartDate *time.Time
    FinishDate *time.Time
    Rated int
    Score int
    Dvd int
    Storage *StorageType
    Status *ReviewStatus
    Comments string
    TimesWatched int
    RewatchValue *RewatchValue
    Tags string
    Rewatching int
    RewatchingEp int
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
            storage_type_id,
            user_status_id,
            comments,
            times_watched,
            rewatch_value,
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
    var storageTypeId int
    var userStatusId int

    row.Scan(
        &review.Id, &review.WatchedEpisodes, &review.StartDate,
        &review.FinishDate, &review.Rated, &review.Score, &review.Dvd,
        &storageTypeId, &userStatusId, &review.Comments, &review.TimesWatched,
        &review.RewatchValue, &review.Tags, &review.Rewatching, &review.RewatchingEp,
    )

    review.Series = GetSeries(db, seriesId)

    return &review
}
