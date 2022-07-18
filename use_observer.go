package main

import (
	"fmt"

	"github.com/luowensheng/design_patterns/observer"
)

func viewObserver(){
	fmt.Println("_________________________________\n") 	

    fmt.Println("Starting Observer") 	
    fmt.Println("_________________________________\n") 	
	
	var dataSubject observer.Observable = observer.NewDataSubject("hello World")

	listeners := []observer.Observer {}


	for i:=0; i< 9; i++ {
         listeners = append(listeners, &observer.DataListener{Id: i})
	}

	for _, listener := range listeners {
		dataSubject.RegisterObserver(listener)
	}
	println("The Number of Listeners (START): ", dataSubject.GetNumberOfListeners())

	dataSubject.NotifyAll()

	for _, listener := range listeners {
		dataSubject.UnregisterObserver(listener)
	}

	println("The Number of Listeners (END): ", dataSubject.GetNumberOfListeners())

}