package adapter

type Televison interface {
	VolumeUp() int
	VolumeDown() int
	ChannelUp() int
	ChannelDown() int
    TurnOn()
	TurnOff()
	GoToChannel(channel int)
}