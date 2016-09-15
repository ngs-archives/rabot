package app

import "encoding/xml"

type Station struct {
	XMLName   xml.Name `xml:"station"`
	ID        string   `xml:"id"`
	Name      string   `xml:"name"`
	AsciiName string   `xml:"ascii_name"`
	Href      string   `xml:"href"`
}

type Stations struct {
	XMLName  xml.Name  `xml:"stations"`
	Stations []Station `xml:"station"`
}

func GetStations(xmlDoc string) []Station {
	res := Stations{}
	xml.Unmarshal([]byte(xmlDoc), &res)
	return res.Stations
}
