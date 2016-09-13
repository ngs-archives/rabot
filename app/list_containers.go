package app

import (
	"bytes"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/stringid"
	"github.com/docker/docker/pkg/stringutils"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/net/context"
	"strings"
	"time"
)

func stripNamePrefix(ss []string) []string {
	sss := make([]string, len(ss))
	for i, s := range ss {
		sss[i] = s[1:]
	}
	return sss
}

func (app *App) ListContainers() string {
	list, err := app.Docker.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return err.Error()
	}
	if len(list) == 0 {
		return "No containers are running"
	}
	var buf bytes.Buffer
	table := tablewriter.NewWriter(&buf)
	table.SetHeader([]string{"CONTAINER ID", "COMMAND", "CREATED", "STATUS", "NAME"})
	for _, v := range list {
		table.Append([]string{stringid.TruncateID(v.ID),
			stringutils.Ellipsis(v.Command, 20),
			time.Unix(int64(v.Created), 0).String(),
			v.Status,
			strings.Join(stripNamePrefix(v.Names), ",")})
	}
	table.SetBorder(false)
	table.Render()
	return "\n```\n" + buf.String() + "```\n"
}
