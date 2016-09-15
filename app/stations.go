package app

import (
	"bytes"
	"encoding/xml"
	"github.com/olekukonko/tablewriter"
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

func (app *App) ListStations(prefIdOrName string) string {
	list, err := FetchStations(prefIdOrName)
	if err != nil {
		return err.Error()
	}
	var buf bytes.Buffer
	table := tablewriter.NewWriter(&buf)
	table.SetHeader([]string{"ID", "NAME", "URL"})
	for _, v := range list {
		table.Append([]string{v.ID, v.Name, v.Href})
	}
	table.SetBorder(false)
	table.Render()
	return "\n```\n" + buf.String() + "```\n"
}
