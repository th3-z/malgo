package parsers

type MalXml struct {
    UserXml *UserXml
    AnimeXml *[]AnimeXml
}


func ParseMalXml(xml string) *MalXml {
    malXml := MalXml {
        UserXml: ParseUserXml(xml),
        AnimeXml: ParseAnimeXml(xml),
    }

    return &malXml
}

