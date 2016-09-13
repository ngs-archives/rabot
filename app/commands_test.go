package app

import (
	"fmt"
	"testing"
)

func TestBuildSetBotID(t *testing.T) {
	app = &App{}
	app.SetBotID("rabot-test")
  t.Fail()
	fmt.Printf("%+v\n%+v\n%+v\n%+v\n", app.Commands.List, app.Commands.Start, app.Commands.Remove, app.Commands.Ping)
}
