package app

import (
	"regexp"
)

type Commands struct {
	List        *regexp.Regexp
	Start       *regexp.Regexp
	Remove      *regexp.Regexp
	Ping        *regexp.Regexp
	Prefectures *regexp.Regexp
	Stations    *regexp.Regexp
}

func (commands *Commands) SetBotID(botID string) {
	commands.List = regexp.MustCompile(`\A<@` + botID + `>\s+(?:list|ls)(?:\s+containers?)?\s*$`)
	commands.Start = regexp.MustCompile(`\A<@` + botID + `>\s+start\s+record(?:ing)?\s+(\S+)\s+(?:for\s+)?(\d+)\s*min(?:utes?)?\s*$`)
	commands.Remove = regexp.MustCompile(`\A<@` + botID + `>\s+(?:remove|rm)\s+(?:container\s+)?(\S+)\s*$`)
	commands.Ping = regexp.MustCompile(`\A<@` + botID + `>\s+ping\s*$`)
	commands.Prefectures = regexp.MustCompile(`\A<@` + botID + `>\s+(?:list|ls)\s+pref(?:ecture)?s?\s*$`)
	commands.Stations = regexp.MustCompile(`\A<@` + botID + `>\s+(?:list|ls)\s+stationss?\s+(\S+)\s*$`)
}
