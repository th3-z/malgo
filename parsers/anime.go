package parsers

import (
	"github.com/antchfx/xmlquery"
	"strconv"
	"strings"
)

type AnimeXml struct {
	SeriesAnimedbId int
	SeriesTitle     string
	SeriesType      string
	SeriesEpisodes  int
	// <my_id> ignored
	MyId              int
	MyWatchedEpisodes int
	MyStartDate       string
	MyFinishDate      string
	MyRated           int
	MyScore           int
	MyDvd             int
	MyStorage         string
	MyStatus          string
	MyComments        string
	MyTimesWatched    int
	MyRewatchValue    string
	MyTags            string
	MyRewatching      int
	MyRewatchingEp    int
	// <update_on_import> ignored
}

func ParseAnimeXml(xml string) *[]AnimeXml {
	animeXml := make([]AnimeXml, 0, 0)

	doc, err := xmlquery.Parse(strings.NewReader(xml))
	if err != nil {
		panic(err)
	}

	animePath := "//anime"
	animeTree := xmlquery.Find(doc, animePath)

	for _, animeTreeV := range animeTree {
		seriesTitle := animeTreeV.SelectElement("series_title").InnerText()
		seriesType := animeTreeV.SelectElement("series_type").InnerText()
		myStartDate := animeTreeV.SelectElement("my_start_date").InnerText()
		myFinishDate := animeTreeV.SelectElement("my_finish_date").InnerText()
		myStorage := animeTreeV.SelectElement("my_storage").InnerText()
		myStatus := animeTreeV.SelectElement("my_status").InnerText()
		myComments := animeTreeV.SelectElement("my_comments").InnerText()
		myTags := animeTreeV.SelectElement("my_tags").InnerText()
		myRewatchValue := animeTreeV.SelectElement("my_rewatch_value").InnerText()

		seriesAnimedbId, _ := strconv.Atoi(animeTreeV.SelectElement("series_animedb_id").InnerText())
		seriesEpisodes, _ := strconv.Atoi(animeTreeV.SelectElement("series_episodes").InnerText())
		myWatchedEpisodes, _ := strconv.Atoi(animeTreeV.SelectElement("my_watched_episodes").InnerText())
		myId, _ := strconv.Atoi(animeTreeV.SelectElement("my_id").InnerText())
		myScore, _ := strconv.Atoi(animeTreeV.SelectElement("my_score").InnerText())
		myRated, _ := strconv.Atoi(animeTreeV.SelectElement("my_rated").InnerText())
		myDvd, _ := strconv.Atoi(animeTreeV.SelectElement("my_dvd").InnerText())
		myTimesWatched, _ := strconv.Atoi(animeTreeV.SelectElement("my_times_watched").InnerText())
		myRewatching, _ := strconv.Atoi(animeTreeV.SelectElement("my_rewatching").InnerText())
		myRewatchingEp, _ := strconv.Atoi(animeTreeV.SelectElement("my_rewatching_ep").InnerText())

		a := AnimeXml{
			SeriesTitle:       seriesTitle,
			SeriesType:        seriesType,
			SeriesAnimedbId:   seriesAnimedbId,
			SeriesEpisodes:    seriesEpisodes,
			MyId:              myId,
			MyWatchedEpisodes: myWatchedEpisodes,
			MyStartDate:       myStartDate,
			MyFinishDate:      myFinishDate,
			MyScore:           myScore,
			MyRated:           myRated,
			MyDvd:             myDvd,
			MyStorage:         myStorage,
			MyStatus:          myStatus,
			MyComments:        myComments,
			MyTags:            myTags,
			MyTimesWatched:    myTimesWatched,
			MyRewatchValue:    myRewatchValue,
			MyRewatchingEp:    myRewatchingEp,
			MyRewatching:      myRewatching,
		}
		animeXml = append(animeXml, a)
	}

	return &animeXml
}
