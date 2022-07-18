package main

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/luowensheng/design_patterns/singleton"
)

func viewSingleton(){
	
	fmt.Println("_________________________________\n") 	

    fmt.Println("Starting Singleton") 	
    fmt.Println("_________________________________\n") 	
	
	var logger = singleton.GetLoggerInstance()
	logger.SetLogLevel(1)
	logger.Log("this is a log message")

	logger = singleton.GetLoggerInstance()
	logger.SetLogLevel(2)
	logger.Log("this is another log message")
    
	var wg sync.WaitGroup

	for i:=0; i< 10; i++ {

		wg.Add(1) 
        
		go func(i int){
			log := singleton.GetLoggerInstance()
			log.Log("message: "+strconv.Itoa(i))
			wg.Done()
			
		}(i)

	}
	wg.Wait()
}

