package malgo

import (
	"database/sql"
	"github.com/th3-z/malgo/parsers"
	"github.com/th3-z/malgo/models"
	"io/ioutil"
)

func MigrateFile(db *sql.DB, filename string) {
	xmlBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	MigrateString(db, string(xmlBytes))
}

func MigrateString(db *sql.DB, xml string) {
	// Inserts spawn transactions, they are slow, run this in one transaction
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Commit()

	malXml := parsers.ParseMalXml(xml)
    user := models.NewUser(db, malXml.UserXml.UserName)
    print(user.Name)
	print("\n")

	for _, animeXml := range *malXml.AnimeXml {
        series := models.NewSeries(db, animeXml.SeriesTitle)
        series.AnimedbId = animeXml.SeriesAnimedbId
        series.Episodes = animeXml.SeriesEpisodes

        review := models.NewReview(db, user.Id, series.Id)
        review.WatchedEpisodes = animeXml.MyWatchedEpisodes
        review.StartDate = animeXml.MyStartDate
        review.FinishDate = animeXml.MyFinishDate
        review.Rated = animeXml.MyRated
        review.Score = animeXml.MyScore
        // review.Dvd = animeXml.MyDvd
        review.Storage = animeXml.MyStorage
        review.Status = animeXml.MyStatus
        review.Comments = animeXml.MyComments
        review.TimesWatched = animeXml.MyTimesWatched
        review.RewatchValue = animeXml.MyRewatchValue
        review.Tags = animeXml.MyTags
        review.Rewatching = animeXml.MyRewatching
        review.RewatchingEp = animeXml.MyRewatchingEp

        user.Reviews = append(user.Reviews, review)

		// models.AddSeries(tx, &anime)
		// models.AddUserAnime(tx, &user, anime.SeriesId)
		// statusId := models.AddUserStatus(tx, &user, anime.SeriesStatus)
		// models.SetAnimeUserStatus(tx, &user, &anime, statusId)
	}
}

