package main

import (
	"fmt"
	"github.com/ngs/rabot/app"
	"github.com/nlopes/slack"
)

func main() {
	app := app.New()
	go app.ManageConnection()

Loop:
	for {
		select {
		case msg := <-app.RTM.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				app.SetBotID(ev.Info.User.ID)
			case *slack.MessageEvent:
				app.HandleMessage(ev)
			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())
			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop
			}
		}
	}
}
