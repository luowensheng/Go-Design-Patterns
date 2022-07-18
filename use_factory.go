package main

import (
	"log"
	"github.com/luowensheng/design_patterns/factory"
	"fmt"
)

func viewFactory(){
	fmt.Println("_________________________________\n") 	

    fmt.Println("Starting Factory") 	
    fmt.Println("_________________________________\n") 	
	
	var pub, err = factory.NewPublication(50,"newspaper" ,"Harry Potter", "New") 
	
	if err != nil {
		log.Fatal(err)
	}
	 pub.SetName("Harry's Potter")
	fmt.Println(pub)
}
