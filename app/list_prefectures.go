package app

import (
	"bytes"
	"github.com/olekukonko/tablewriter"
)

func (app *App) ListPrefectures() string {
	var buf bytes.Buffer
	table := tablewriter.NewWriter(&buf)
	table.SetHeader([]string{"ID", "PREF"})
	for _, v := range Prefectures {
		table.Append([]string{v.ID, v.Name})
	}
	table.SetBorder(false)
	table.Render()
	return "\n```\n" + buf.String() + "```\n"
}
