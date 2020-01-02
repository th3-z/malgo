package path

import(
    _"github.com/antchfx/xpath"
    "github.com/antchfx/xmlquery"
    "strings"
    "strconv"
)


type user struct {
    Name string
    MalUserId int  // Unused
    ExportType int  // Unused
}

type anime struct {
    Name string
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

    malUserId, _ := strconv.Atoi(userTree.SelectElement("user_id").InnerText())
    exportType, _ := strconv.Atoi(userTree.SelectElement("user_export_type").InnerText())

    u := user{
        Name: userTree.SelectElement("user_name").InnerText(),
        MalUserId: malUserId,
        ExportType: exportType,
    }

    return &u
}

