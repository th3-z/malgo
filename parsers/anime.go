package parsers

import (
	"github.com/antchfx/xmlquery"
	"strconv"
	"strings"
)

type AnimeXml struct {
	SeriesAnimedbId   int
	SeriesTitle       string
	SeriesType        string
	SeriesEpisodes    int
	MyId              int     // Unused
	MyWatchedEpisodes int
	MyStartDate       int
	MyFinishDate      int
	MyRated           int     // Unused
	MyScore           int
	MyDvd             string  // Unused
    MyStorage         string
	MyStatus          string
	MyComments        string
	MyTimesWatched    int
	MyRewatchValue    int
	MyTags            string
	MyRewatching      int
	MyRewatchingEp    int
	UpdateOnImport    int     // Unused
}

func ParseAnimeXml(xml string) *[]AnimeXml {
	animeXml := make([]AnimeXml, 0, 0)

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
		seriesAnimedbId, _ := strconv.Atoi(animeTreeV.SelectElement("series_animedb_id").InnerText())
		seriesEpisodes, _ := strconv.Atoi(animeTreeV.SelectElement("series_episodes").InnerText())
		myWatchedEpisodes, _ := strconv.Atoi(animeTreeV.SelectElement("my_watched_episodes").InnerText())
		myStartDate, _ := strconv.Atoi(animeTreeV.SelectElement("my_start_date").InnerText())
		myFinishDate, _ := strconv.Atoi(animeTreeV.SelectElement("my_finish_date").InnerText())
		myScore, _ := strconv.Atoi(animeTreeV.SelectElement("my_score").InnerText())
		myTimesWatched, _ := strconv.Atoi(animeTreeV.SelectElement("my_times_watched").InnerText())
		myRewatchValue, _ := strconv.Atoi(animeTreeV.SelectElement("my_rewatch_value").InnerText())
		myRewatchingEp, _ := strconv.Atoi(animeTreeV.SelectElement("my_rewatching_ep").InnerText())

		a := AnimeXml{
			SeriesTitle:       seriesTitle,
			SeriesType:        seriesType,
			SeriesAnimedbId:   seriesAnimedbId,
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
		animeXml = append(animeXml, a)
	}

	return &animeXml
}

