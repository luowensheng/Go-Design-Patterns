package main

import (
	"fmt"

	"github.com/luowensheng/design_patterns/facade"
)

func viewFacade(){
    fmt.Println("_________________________________\n") 	

    fmt.Println("Starting Facade") 	
    fmt.Println("_________________________________\n") 	
	facade.MakeAmericano(4)
	facade.MakeLatte(40)

}

