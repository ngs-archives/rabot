package main

import (
	"bytes"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stringid"
	"github.com/docker/docker/pkg/stringutils"
	"github.com/nlopes/slack"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/net/context"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
	// "github.com/docker/go-connections/sockets"
)

func main() {
	client, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(os.Stdout, "rabot: ", log.Lshortfile|log.LstdFlags)
	api := slack.New(os.Getenv("SLACK_TOKEN"))
	slack.SetLogger(logger)
	api.SetDebug(os.Getenv("DEBUG") != "")

	botname := os.Getenv("RABOT_NAME")
	if botname == "" {
		botname = "rabot"
	}

	var listCommand = regexp.MustCompile(`\A$`)
	var startCommand = regexp.MustCompile(`\A$`)
	var removeCommand = regexp.MustCompile(`\A$`)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			var reply = ""
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				botID := ev.Info.User.ID
				listCommand = regexp.MustCompile(`\A<@` + botID + `>\s+list\s+containers?`)
				startCommand = regexp.MustCompile(`\A<@` + botID + `>\s+start\s+record(?:ing)?\s+(\S+)\s+(?:for\s+)?(\d+)\s*min(?:utes?)?`)
				removeCommand = regexp.MustCompile(`\A<@` + botID + `>\s+remove\s+(?:container\s+)?(\S+)`)
			case *slack.MessageEvent:
				if listCommand.MatchString(ev.Text) {
					reply = ListContainers(client)
				} else if res := startCommand.FindStringSubmatch(ev.Text); res != nil {
					channel, err := api.GetChannelInfo(ev.Channel)
					if err != nil {
						reply = err.Error()
					} else {
						reply = StartContainer(client, res[1], res[2], channel.Name)
					}
				} else if res := removeCommand.FindStringSubmatch(ev.Text); res != nil {
					reply = RemoveContainer(client, res[1])
				}
				if reply != "" {
					rtm.SendMessage(rtm.NewOutgoingMessage("<@"+ev.User+"> "+reply, ev.Channel))
				}
			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())
			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop
			}
		}
	}
}

func stripNamePrefix(ss []string) []string {
	sss := make([]string, len(ss))
	for i, s := range ss {
		sss[i] = s[1:]
	}
	return sss
}

func ListContainers(client *client.Client) string {
	list, err := client.ContainerList(context.Background(), types.ContainerListOptions{})
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

func StartContainer(client *client.Client, station string, minutes string, channel string) string {
	imageName := os.Getenv("IMAGE_NAME")
	if imageName == "" {
		imageName = "atsnngs/radiko-recorder-s3"
	}
	config := &container.Config{
		Image:        imageName,
		AttachStdin:  false,
		AttachStdout: false,
		AttachStderr: false,
		StdinOnce:    false,
		Env: []string{"SLACK_CHANNEL=" + channel,
			"STATION=" + station,
			"DURATION_MINUTES=" + minutes,
			"RADIKO_LOGIN=" + os.Getenv("RADIKO_LOGIN"),
			"RADIKO_PASSWORD=" + os.Getenv("RADIKO_PASSWORD"),
			"S3_BUCKET=" + os.Getenv("S3_BUCKET"),
			"AWS_ACCESS_KEY_ID=" + os.Getenv("AWS_ACCESS_KEY_ID"),
			"AWS_SECRET_ACCESS_KEY=" + os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"SLACK_WEBHOOK_URL=" + os.Getenv("SLACK_WEBHOOK_URL")}}
	hostConfig := &container.HostConfig{
		AutoRemove: true}
	ctx := context.Background()
	res, err := client.ContainerCreate(ctx, config, hostConfig, nil, "")
	if err != nil {
		return err.Error()
	}
	if err := client.ContainerStart(ctx, res.ID, types.ContainerStartOptions{}); err != nil {
		return err.Error()
	}
	return "Started container " + stringid.TruncateID(res.ID)
}

func RemoveContainer(client *client.Client, containerID string) string {
	ctx := context.Background()
	if err := client.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{Force: true}); err != nil {
		return err.Error()
	}
	return "Successfully removed " + containerID
}
