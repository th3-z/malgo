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
    user := models.NewUser(tx, malXml.UserXml.UserName)
    print(user.Name)
	print("\n")

	for _, animeXml := range *malXml.AnimeXml {
        series := models.NewSeries(tx, animeXml.SeriesTitle)
		series.Type = models.NewSeriesType(tx, animeXml.SeriesType)
        series.AnimedbId = animeXml.SeriesAnimedbId
        series.Episodes = animeXml.SeriesEpisodes

        review := models.NewReview(tx, user.Id, series.Id)
        review.Status = models.NewUserStatus(tx, animeXml.MyStatus)
        review.WatchedEpisodes = animeXml.MyWatchedEpisodes
        // review.StartDate = animeXml.MyStartDate TODO
        // review.FinishDate = animeXml.MyFinishDate TODO
        review.Rated = animeXml.MyRated
        review.Score = animeXml.MyScore
        review.Dvd = animeXml.MyDvd
        review.Storage = models.NewStorageType(tx, animeXml.MyStorage)
        review.Comments = animeXml.MyComments
        review.TimesWatched = animeXml.MyTimesWatched
        review.RewatchValue = models.NewRewatchValue(tx, animeXml.MyRewatchValue)
        review.Tags = animeXml.MyTags
        review.Rewatching = animeXml.MyRewatching
        review.RewatchingEp = animeXml.MyRewatchingEp

        user.Reviews = append(user.Reviews, review)

        print(series.Title)
        print(" -> ")
        print(series.Type.Id)

        print("\n")
	}
}
