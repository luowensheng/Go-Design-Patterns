package main

import (
	"log"
	"github.com/luowensheng/design_patterns/factory"
	"fmt"
)

func viewFactory(){
	
	var pub, err = factory.NewPublication(50,"newspaper" ,"Harry Potter", "New") 
	
	if err != nil {
		log.Fatal(err)
	}
	 pub.SetName("Harry's Potter")
	fmt.Println(pub)
}
