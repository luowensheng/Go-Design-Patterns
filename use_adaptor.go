package main

import (
	"fmt"

	"github.com/luowensheng/design_patterns/adapter"
)

func viewAdapter() {

	fmt.Println("_________________________________\n") 	

    fmt.Println("Starting Adapter") 	
    fmt.Println("_________________________________\n") 	
	
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
