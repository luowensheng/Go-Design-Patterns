package adapter

type SammysangTVAdapter struct {
	SammysangTV
}

func (tv *SammysangTVAdapter) VolumeUp() int {
	tv.setVolume(tv.currentVolume + 1)
	return tv.getVolume()
}

func (tv *SammysangTVAdapter) VolumeDown() int {
	tv.setVolume(tv.currentVolume - 1)
	return tv.getVolume()

}

func (tv *SammysangTVAdapter) ChannelUp() int {
	tv.setChannel(tv.currentChan + 1)
	return tv.getChannel()
}

func (tv *SammysangTVAdapter) ChannelDown() int {
	tv.setChannel(tv.currentChan - 1)
	return tv.getChannel()
}

func (tv *SammysangTVAdapter) TurnOn() {
	tv.setOnState(true)
}

func (tv *SammysangTVAdapter) TurnOff() {
	tv.setOnState(false)
}

func (tv *SammysangTVAdapter) GoToChannel(channel int) {
	tv.setChannel(channel)

}
