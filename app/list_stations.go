package app

import (
	"bytes"
	"github.com/olekukonko/tablewriter"
)

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
