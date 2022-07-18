package main

import (
	"github.com/luowensheng/design_patterns/adapter"
)

func viewAdapter() {

	tvs := []adapter.Televison{
		&adapter.SonheeTV{},
		&adapter.SammysangTVAdapter{},
	}

	for _, tv := range tvs {
		tv.TurnOn()
		tv.VolumeUp()
		tv.VolumeDown()
		tv.ChannelUp()
		tv.ChannelDown()
		tv.GoToChannel(69)
		tv.TurnOff()
	}

}
