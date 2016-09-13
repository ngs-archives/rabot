package app

import (
	dockerClient "github.com/docker/docker/client"
	"github.com/nlopes/slack"
	"log"
	"os"
)

type App struct {
	Slack  *slack.Client
	Docker *dockerClient.Client
	BotID  string
	RTM    *slack.RTM
	Commands
}

func New() *App {
	dockerClient, err := dockerClient.NewEnvClient()
	if err != nil {
		panic(err)
	}
	logger := log.New(os.Stdout, "rabot: ", log.Lshortfile|log.LstdFlags)
	slackClient := slack.New(os.Getenv("SLACK_TOKEN"))
	slack.SetLogger(logger)
	slackClient.SetDebug(os.Getenv("DEBUG") != "")
	RTM := slackClient.NewRTM()
	return &App{
		Slack:  slackClient,
		Docker: dockerClient,
		RTM:    RTM,
	}
}

func (app *App) SetBotID(botID string) {
	app.BotID = botID
	app.Commands.SetBotID(botID)
}

func (app *App) ManageConnection() {
	app.RTM.ManageConnection()
}

func (app *App) BuildReplyMessage(ev *slack.MessageEvent) string {
	text := ev.Text
	if app.Commands.List.MatchString(text) {
		return app.ListContainers()
	}
	if res := app.Commands.Start.FindStringSubmatch(text); res != nil {
		channel, err := app.Slack.GetChannelInfo(ev.Channel)
		if err != nil {
			return err.Error()
		} else {
			return app.StartContainer(res[1], res[2], channel.Name)
		}
	} else if res := app.Commands.Remove.FindStringSubmatch(text); res != nil {
		return app.RemoveContainer(res[1])
	}
	return ""
}

func (app *App) HandleMessage(ev *slack.MessageEvent) {
	if reply := app.BuildReplyMessage(ev); reply != "" {
		app.RTM.SendMessage(app.RTM.NewOutgoingMessage("<@"+ev.User+"> "+reply, ev.Channel))
	}
}
