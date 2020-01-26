package parsers

import (
	"github.com/antchfx/xmlquery"
	"strconv"
	"strings"
)

type UserXml struct {
	UserId   int
	UserName string
	// <user_export_type> ignored
	// <user_total_anime> ignored
	// <user_total_watching> ignored
	// <user_total_completed> ignored
	// <user_total_onhold> ignored
	// <user_total_dropped> ignored
	// <user_total_plantowatch> ignored
}

func ParseUserXml(xml string) *UserXml {
	doc, err := xmlquery.Parse(strings.NewReader(xml))
	if err != nil {
		panic(err)
	}

	userPath := "//myinfo"
	userTree := xmlquery.FindOne(doc, userPath)

	userId, _ := strconv.Atoi(userTree.SelectElement("user_id").InnerText())
	userName := userTree.SelectElement("user_name").InnerText()

	userXml := UserXml{
		UserId:   userId,
		UserName: userName,
	}

	return &userXml
}
