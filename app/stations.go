package app

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

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

func GetStations(xmlDoc []byte) []Station {
	res := Stations{}
	xml.Unmarshal(xmlDoc, &res)
	return res.Stations
}

func FetchStations(prefIdOrName string) ([]Station, error) {
	pref, err := FindPrefecture(prefIdOrName)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get("http://radiko.jp/v2/station/list/" + pref.ID + ".xml")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	stations := GetStations(body)
	return stations, nil
}
