package main

import (
	"bytes"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
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
	channel := os.Getenv("SLACK_CHANNEL")

	listCommand := regexp.MustCompile(`\A` + botname + `\s+list\s+containers?`)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				if listCommand.MatchString(ev.Text) {
					table := ListContainers(client)
					fmt.Printf("%v %v", channel, table)
					rtm.SendMessage(rtm.NewOutgoingMessage(table, "#"+channel))
				} else {
					fmt.Printf("Message: %v %v\n", ev.Text)
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

	var buf bytes.Buffer
	table := tablewriter.NewWriter(&buf)
	table.SetHeader([]string{"Name", "CREATED", "STATUS"})
	for _, v := range list {
		table.Append([]string{strings.Join(stripNamePrefix(v.Names), ","), time.Unix(int64(v.Created), 0).String(), v.Status})
	}
	table.SetBorder(false)
	table.Render()
	return buf.String()
}
