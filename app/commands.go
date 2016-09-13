package app

import (
	"regexp"
)

type Commands struct {
	List   *regexp.Regexp
	Start  *regexp.Regexp
	Remove *regexp.Regexp
	Ping   *regexp.Regexp
}

func (commands *Commands) SetBotID(botID string) {
	commands.List = regexp.MustCompile(`\A<@` + botID + `>\s+(?:list|ls)(?:\s+containers?)?`)
	commands.Start = regexp.MustCompile(`\A<@` + botID + `>\s+start\s+record(?:ing)?\s+(\S+)\s+(?:for\s+)?(\d+)\s*min(?:utes?)?`)
	commands.Remove = regexp.MustCompile(`\A<@` + botID + `>\s+(?:remove|rm)\s+(?:container\s+)?(\S+)`)
	commands.Ping = regexp.MustCompile(`\A<@` + botID + `>\s+ping`)
}
