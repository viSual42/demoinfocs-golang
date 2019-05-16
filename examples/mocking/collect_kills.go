package mocking

import (
	dem "github.com/visual42/demoinfocs-golang"
	events "github.com/visual42/demoinfocs-golang/events"
)

func collectKills(parser dem.IParser) (kills []events.Kill, err error) {
	parser.RegisterEventHandler(func(kill events.Kill) {
		kills = append(kills, kill)
	})
	err = parser.ParseToEnd()
	return
}
