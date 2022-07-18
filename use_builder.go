package main

import (
	"fmt"
	"log"
	"github.com/luowensheng/design_patterns/builder"
)

func viewBuilder(){
	
	fmt.Println("_________________________________\n") 	

    fmt.Println("Starting Builder") 	
    fmt.Println("_________________________________\n") 	

	var nb = builder.NewNotificationBuilder()
    
	nb.SetTitle("New Notification")
	nb.SetMessage("This is a basic Notification")
	nb.SetIcon("icon.png")
	nb.SetSubTitle("icon")
	nb.SetImage("image.png")
	nb.SetPriority(2)
	nb.SetNotType("already")


	notification, err := nb.Build()
    
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(notification)
}