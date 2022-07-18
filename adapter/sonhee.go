package adapter

import "fmt"

type SonheeTV struct {
	vol int
	channel int
	isOn bool
}

func (tv *SonheeTV) VolumeUp() int {
   tv.vol ++
   fmt.Println("SonheeTV vol", tv.vol)
   return tv.vol
}

func (tv *SonheeTV) VolumeDown() int {
	tv.vol --
  fmt.Println("SonheeTV vol", tv.vol)
	return tv.vol

}

func (tv *SonheeTV) ChannelUp() int {
    tv.channel++
    fmt.Println("SonheeTV channel", tv.channel)
    return tv.channel
}

func (tv *SonheeTV) ChannelDown() int {
   tv.channel --
   fmt.Println("SonheeTV channel", tv.channel)
   return tv.channel
}

func (tv *SonheeTV) TurnOn(){
  tv.isOn = true
  fmt.Println("SonheeTV isOn", tv.isOn)
}

func (tv *SonheeTV) TurnOff(){
  tv.isOn = false
  fmt.Println("SonheeTV isOn", tv.isOn)
}

func (tv *SonheeTV) GoToChannel(channel int){
  tv.channel = channel
  fmt.Println("SonheeTV channel", tv.channel)
}