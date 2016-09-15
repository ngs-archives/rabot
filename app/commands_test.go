package app

import (
	"regexp"
	"testing"
)

func _TestCommands(t *testing.T, regexp *regexp.Regexp, commands []string) {
	for _, cmd := range commands {
		if !regexp.MatchString(cmd) {
			t.Errorf("`%v` should match with `%v`, but not", cmd, regexp)
		}
	}
}

func TestListCommand(t *testing.T) {
	commands := &Commands{}
	commands.SetBotID("rabot-test")
	_TestCommands(t, commands.List, []string{
		"<@rabot-test>  list   container  ",
		"<@rabot-test>  list   containers  ",
		"<@rabot-test>  ls  containers ",
		"<@rabot-test>  ls ",
	})
}

func TestStartCommand(t *testing.T) {
	commands := &Commands{}
	commands.SetBotID("rabot-test")
	_TestCommands(t, commands.Start, []string{
		"<@rabot-test>  start  record    ALPHA-STATION for  12 min ",
		"<@rabot-test>  start  recording    ALPHA-STATION   1minute ",
	})
}

func TestRemoveCommand(t *testing.T) {
	commands := &Commands{}
	commands.SetBotID("rabot-test")
	_TestCommands(t, commands.Remove, []string{
		"<@rabot-test>  remove  container  foo  ",
		"<@rabot-test>  rm  container  foo  ",
		"<@rabot-test>  rm   foo  ",
	})
}

func TestPingCommand(t *testing.T) {
	commands := &Commands{}
	commands.SetBotID("rabot-test")
	_TestCommands(t, commands.Ping, []string{
		"<@rabot-test>  ping  ",
	})
}

func TestPrefecturesCommand(t *testing.T) {
	commands := &Commands{}
	commands.SetBotID("rabot-test")
	_TestCommands(t, commands.Prefectures, []string{
		"<@rabot-test>  list   prefectures  ",
		"<@rabot-test>  list   prefecture  ",
		"<@rabot-test>  ls   prefectures  ",
		"<@rabot-test>  ls   prefecture  ",
		"<@rabot-test>  list   prefs  ",
		"<@rabot-test>  list   pref  ",
		"<@rabot-test>  ls   prefs  ",
		"<@rabot-test>  ls   pref  ",
	})
}

func TestStationsCommand(t *testing.T) {
	commands := &Commands{}
	commands.SetBotID("rabot-test")
	_TestCommands(t, commands.Stations, []string{
		"<@rabot-test>  list   stations  京都   ",
		"<@rabot-test>  ls   stations  京都   ",
	})
}
