package parsers

import (
	"github.com/antchfx/xmlquery"
	_ "github.com/antchfx/xpath"
	"strconv"
	"strings"
    "mal-sqlite-migrate/models"
)

func ParseAnimeList(xml string) []models.Anime {
	animeList := make([]models.Anime, 0, 0)

	doc, err := xmlquery.Parse(strings.NewReader(xml))
	if err != nil {
		panic(err)
	}

	animePath := "//anime"

	animeTree := xmlquery.Find(doc, animePath)
	if err != nil {
		panic(err)
	}

	for _, animeTreeV := range animeTree {
		seriesTitle := animeTreeV.SelectElement("series_title").InnerText()
		seriesType := animeTreeV.SelectElement("series_type").InnerText()
		myStorage := animeTreeV.SelectElement("my_storage").InnerText()
		myStatus := animeTreeV.SelectElement("my_status").InnerText()
		myComments := animeTreeV.SelectElement("my_comments").InnerText()
		myTags := animeTreeV.SelectElement("my_tags").InnerText()
		seriesAnimeDbId, _ := strconv.Atoi(animeTreeV.SelectElement("series_animedb_id").InnerText())
		seriesEpisodes, _ := strconv.Atoi(animeTreeV.SelectElement("series_episodes").InnerText())
		myWatchedEpisodes, _ := strconv.Atoi(animeTreeV.SelectElement("my_watched_episodes").InnerText())
		myStartDate, _ := strconv.Atoi(animeTreeV.SelectElement("my_start_date").InnerText())
		myFinishDate, _ := strconv.Atoi(animeTreeV.SelectElement("my_finish_date").InnerText())
		myScore, _ := strconv.Atoi(animeTreeV.SelectElement("my_score").InnerText())
		myTimesWatched, _ := strconv.Atoi(animeTreeV.SelectElement("my_times_watched").InnerText())
		myRewatchValue, _ := strconv.Atoi(animeTreeV.SelectElement("my_rewatch_value").InnerText())
		myRewatchingEp, _ := strconv.Atoi(animeTreeV.SelectElement("my_rewatching_ep").InnerText())

		a := models.Anime{
			SeriesTitle:       seriesTitle,
			SeriesType:        seriesType,
			SeriesAnimeDbId:   seriesAnimeDbId,
			SeriesEpisodes:    seriesEpisodes,
			MyWatchedEpisodes: myWatchedEpisodes,
			MyStartDate:       myStartDate,
			MyFinishDate:      myFinishDate,
			MyScore:           myScore,
			MyStorage:         myStorage,
			MyStatus:          myStatus,
			MyComments:        myComments,
			MyTags:            myTags,
			MyTimesWatched:    myTimesWatched,
			MyRewatchValue:    myRewatchValue,
			MyRewatchingEp:    myRewatchingEp,
		}
		animeList = append(animeList, a)
	}

	return animeList
}
