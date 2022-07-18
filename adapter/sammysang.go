package adapter

import (
	"fmt"
)

type SammysangTV struct {
	currentChan int
	currentVolume int
	tvOn bool
}




func (tv *SammysangTV) getVolume() int {
    return tv.currentVolume
}

func (tv *SammysangTV) setVolume(volume int) {
     tv.currentVolume = volume
	 fmt.Println("SammysangTV volume ", tv.currentVolume)
}

func (tv *SammysangTV) getChannel() int {
    return tv.currentChan
}

func (tv *SammysangTV) setChannel(channel int) {
     tv.currentChan = channel
	 fmt.Println("SammysangTV Channel ", tv.currentChan)
	}

func (tv *SammysangTV) setOnState(tvOn bool) {
    
	tv.tvOn = tvOn

	if tvOn {
		fmt.Println("SammysangTV is ON")
	} else {
		fmt.Println("SammysangTV is OFF")

	}
    
}
