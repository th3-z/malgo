package parsers

import (
	"github.com/antchfx/xmlquery"
	_ "github.com/antchfx/xpath"
	"strconv"
	"strings"
    "mal-sqlite-migrate/models"
)


func ParseUser(xml string) *models.User {
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

	u := models.User{
		UserName:       userTree.SelectElement("user_name").InnerText(),
		UserId:         UserId,
		UserExportType: UserExportType,
	}

	return &u
}

