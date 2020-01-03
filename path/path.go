package path

import (
	"github.com/antchfx/xmlquery"
	_ "github.com/antchfx/xpath"
	"strconv"
	"strings"
)

type user struct {
	UserId               int
	UserName             string
	UserExportType       int // Unused
	UserTotalAnime       int // Unused
	UserTotalWatching    int // Unused
	UserTotalCompleted   int // Unused
	UserTotalOnHold      int // Unused
	UserTotalDropped     int // Unused
	UserTotalPlanToWatch int // Unused
}

type anime struct {
	SeriesTitle       string
	SeriesAnimeDbId   int
	SeriesType        string
	SeriesEpisodes    int
	MyWatchedEpisodes int
	MyStartDate       int
	MyFinishDate      int
	MyScore           int
	MyStorage         string
	MyStatus          string
	MyComments        string
	MyTimesWatched    int
	MyRewatchValue    int
	MyTags            string
	MyRewatching      int
	MyRewatchingEp    int

	MyId           int    // Unused
	MyDvd          string // Unused
	MyRated        int    // Unused
	UpdateOnImport int    // Unused
}

func GetUser(xml string) *user {
	doc, err := xmlquery.Parse(strings.NewReader(xml))
	if err != nil {
		panic(err)
	}

	userPath := "//myinfo"

	userTree := xmlquery.FindOne(doc, userPath)
	if err != nil {
		panic(err)
	}

	UserId, _ := strconv.Atoi(userTree.SelectElement("user_id").InnerText())
	UserExportType, _ := strconv.Atoi(userTree.SelectElement("user_export_type").InnerText())

	u := user{
		UserName:       userTree.SelectElement("user_name").InnerText(),
		UserId:         UserId,
		UserExportType: UserExportType,
	}

	return &u
}

func GetAnimeList(xml string) []anime {
	animeList := make([]anime, 0, 0)

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

		a := anime{
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
