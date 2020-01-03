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
	AnimeDbId         int
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
		SeriesTitle := animeTreeV.SelectElement("series_title").InnerText()
		a := anime{
			SeriesTitle: SeriesTitle,
		}
		animeList = append(animeList, a)
	}

	return animeList
}
