package parsers

import (
	"github.com/antchfx/xmlquery"
	_ "github.com/antchfx/xpath"
	"strconv"
	"strings"
)

type UserXml struct {
    UserId               int  // Unused
    UserName             int
    UserExportType       int  // Unused
    UserTotalAnime       int  // Unused
    UserTotalWatching    int  // Unused
    UserTotalCompleted   int  // Unused
    UserTotalOnhold      int  // Unused
    UserTotalDropped     int  // Unused
    UserTotalPlantowatch int  // Unused
}

func ParseUserXml(xml string) *UserXml {
	doc, err := xmlquery.Parse(strings.NewReader(xml))
	if err != nil {
		panic(err)
	}

	userPath := "//myinfo"

	userTree := xmlquery.FindOne(doc, userPath)
	if err != nil {
		panic(err)
	}

	userId, _ := strconv.Atoi(userTree.SelectElement("user_id").InnerText())
	userName, _ := strconv.Atoi(userTree.SelectElement("user_name").InnerText())

    userXml := UserXml {
        UserId: userId,
        UserName: userName,
    }

    return &userXml
}

